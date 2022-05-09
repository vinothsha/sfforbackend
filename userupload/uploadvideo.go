package userupload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func UploadVideos(w http.ResponseWriter, r *http.Request) {
	var VideoDetails UploadVideo
	VideoDetails.VideoID = uuid.New().String()
	VideoDetails.Createddatetime = time.Now().Format("2006-01-02 15:04:05")
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read video Details from User")
	}
	json.Unmarshal(req, &VideoDetails)
	// Open the file
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("myfile")
	fmt.Println("file info")
	fmt.Println("file Name" + handler.Filename)
	// file, err := os.Open(handler.Filename)
	// file, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) //)os.O_RDWR, 0644 //os.O_RDONLY, 0644
	if err != nil {
		fmt.Println("error in openfile")
	}
	defer file.Close()

	// stats, _ := file.Stat()
	if err != nil {
		fmt.Println("error while find Length of the video")
	}
	fileSize := handler.Size
	FileSizeInMB := (float64)(fileSize / 1024)
	// put file in byteArray
	buffer := make([]byte, fileSize)
	file.Read(buffer)

	// Create MultipartUpload object
	createdResp, err := s3session.CreateMultipartUpload(&s3.CreateMultipartUploadInput{
		Bucket:      aws.String(BUCKET_NAME),
		Key:         aws.String(handler.Filename),
		ACL:         aws.String(s3.BucketCannedACLPublicReadWrite),
		ContentType: aws.String("audio/mpeg"),

		// RequestPayer: aws.String("requester"),
	})

	if err != nil {
		fmt.Println(err)
		fmt.Println("error hit 1")
		return
	}

	var start, currentSize int
	var remaining = int(fileSize)
	var partNum = 1
	var completedParts []*s3.CompletedPart
	// Loop till remaining upload size is 0
	for start = 0; remaining != 0; start += PART_SIZE {
		if remaining < PART_SIZE {
			currentSize = remaining
		} else {
			currentSize = PART_SIZE
		}

		completed, err := Upload(createdResp, buffer[start:start+currentSize], partNum)
		// If upload function failed (meaning it retried acoording to RETRIES)
		if err != nil {
			_, err = s3session.AbortMultipartUpload(&s3.AbortMultipartUploadInput{
				Bucket:   createdResp.Bucket,
				Key:      createdResp.Key,
				UploadId: createdResp.UploadId,
			})
			if err != nil {
				fmt.Println(err)
				fmt.Println("error hit 2")
				return
			}
		}

		// Detract the current part size from remaining
		remaining -= currentSize
		fmt.Printf("Part %v complete, %v btyes remaining\n", partNum, remaining)

		// Add the completed part to our list
		completedParts = append(completedParts, completed)
		partNum++

	}

	// All the parts are uploaded, completing the upload
	resp, err := s3session.CompleteMultipartUpload(&s3.CompleteMultipartUploadInput{
		Bucket:   createdResp.Bucket,
		Key:      createdResp.Key,
		UploadId: createdResp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	})
	if err != nil {
		fmt.Println(err)
		fmt.Println("error hit 3")
	} else {
		// Url := *resp.Location
		fmt.Println(FileSizeInMB / 1024)
		VideoDetails.VideoSizeInMb = FileSizeInMB / 1024
		fmt.Println(*resp.Location)
		fmt.Println(*resp.Key)
		if err := cassession.Session.Query("insert into videos (videouid,videolink,videosizeinmb,title,description,language,geners,agegroup,createddatetime,mail)values(?,?,?,?,?,?,?,?,?,?)",
			VideoDetails.VideoID, *resp.Location, VideoDetails.VideoSizeInMb, VideoDetails.Title, VideoDetails.Description, VideoDetails.Language, VideoDetails.Genres, VideoDetails.AgeGroup, VideoDetails.Createddatetime, VideoDetails.Mail).Exec(); err != nil {
			fmt.Println("error while insert VideoDetails into the Videos table")
		}
		fmt.Println("video uploaded success")
	}
}

// Uploads the fileBytes bytearray a MultiPart upload
func Upload(resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNum int) (completedPart *s3.CompletedPart, err error) {
	var try int
	for try <= RETRIES {
		uploadResp, err := s3session.UploadPart(&s3.UploadPartInput{
			Body:          bytes.NewReader(fileBytes),
			Bucket:        resp.Bucket,
			Key:           resp.Key,
			PartNumber:    aws.Int64(int64(partNum)),
			UploadId:      resp.UploadId,
			ContentLength: aws.Int64(int64(len(fileBytes))),
			// ServerSideEncryption: aws.String("AES256"),
		})
		// Upload failed
		if err != nil {
			fmt.Println(err)
			fmt.Println("error hit 4")
			// Max retries reached! Quitting
			if try == RETRIES {
				fmt.Println("error hit 5")
				return nil, err
			} else {
				// Retrying
				try++
			}
		} else {
			// Upload is done!
			return &s3.CompletedPart{
				ETag:       uploadResp.ETag,
				PartNumber: aws.Int64(int64(partNum)),
			}, nil
		}
	}

	return nil, nil
}
