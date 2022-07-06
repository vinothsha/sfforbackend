package userupload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"
	e "sha/commonservices/commonfunctions"
	s "sha/commonstruct"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

func UploadVideos(w http.ResponseWriter, r *http.Request) {

	var VideoDetails s.UploadVideo
	VideoDetails.UserId = mux.Vars(r)["id"]
	VideoDetails.VideouID = gocql.UUID(uuid.New())
	VideoDetails.Createddatetime = time.Now().Format("2006-01-02 15:04:05")
	r.ParseMultipartForm(32 << 20)
	VideoDetails.AgeGroup = r.Form.Get("agegroup")
	VideoDetails.Title = r.Form.Get("title")
	VideoDetails.Description = r.Form.Get("description")
	VideoDetails.Tags = r.Form.Get("tags")
	VideoDetails.Genres = r.Form.Get("geners")
	// for _, gen := range strings.Split(r.Form.Get("genres"), ",") {
	// 	VideoDetails.Genres = append(VideoDetails.Genres, string(gen))
	// }
	// for _, gen := range strings.Split(r.Form.Get("tags"), ",") {
	// 	VideoDetails.Tags = append(VideoDetails.Tags, string(gen))
	//  }
	VideoDetails.Language = r.Form.Get("language")
	//Get Usermail/Mobile From the Browser Cokkie
	// tokenCookie, _ := r.Cookie("token")
	// a := string(tokenCookie.Value)

	// base64Text := make([]byte, base64.URLEncoding.DecodedLen(len(strings.Split(a, ".")[1])))
	// base64.StdEncoding.Decode(base64Text, []byte(strings.Split(a, ".")[1]))
	// str := string(base64Text)
	// findcolon := strings.Index(str, ":")
	// findcomma := strings.Index(str, ",")
	// VideoDetails.Mail = str[findcolon+2 : findcomma-1]
	// fmt.Println(VideoDetails.Mail)
	//thumnail start
	// Read the entire file into a byte slice

	// fmt.Println(VideoDetails.Mail)
	file1, handler1, err := r.FormFile("myimage")

	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()
	filesize1 := handler1.Size
	buffer1 := make([]byte, filesize1)
	file1.Read(buffer1)
	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(buffer1)
	base64Encoding += "data:" + mimeType + ";base64,"
	// Append the base64 encoded output
	base64Encoding += e.ToBase64(buffer1)
	VideoDetails.Thumnail = base64Encoding
	//end thumnail
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println("error in video openfile")
	}
	defer file.Close()

	if err != nil {
		fmt.Println("error while find Length of the video")
		fmt.Println(err)
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
		// fmt.Printf("Part %v complete, %v btyes remaining\n", partNum, remaining)

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
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Println(err)
		fmt.Println("error hit 3")
	} else {

		VideoDetails.VideoLink = *resp.Location
		VideoDetails.Etag = *resp.ETag
		VideoDetails.VideoSizeInMb = FileSizeInMB / 1024
		// if VideoDetails.Etag != "" {
		var etagfromDb string
		resEtag := cassession.Session.Query("select etag from videos where etag=? allow filtering", VideoDetails.Etag)
		resEtag.Scan(&etagfromDb)

		if etagfromDb != VideoDetails.Etag {
			// res := cassession.Session.Query("select uid from signup where usermail=? allow filtering;", VideoDetails.Mail)
			// res.Scan(&VideoDetails.UserId)
			// fmt.Println(VideoDetails.UserId)

			if err := cassession.Session.Query("insert into videos (videouid,videolink,videosizeinmb,title,description,language,genres,tags,agegroup,createddatetime,useruid,thumnail,etag)values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
				VideoDetails.VideouID, VideoDetails.VideoLink, VideoDetails.VideoSizeInMb, VideoDetails.Title, VideoDetails.Description, VideoDetails.Language, VideoDetails.Genres, VideoDetails.Tags, VideoDetails.AgeGroup, VideoDetails.Createddatetime, VideoDetails.UserId, VideoDetails.Thumnail, VideoDetails.Etag).Exec(); err != nil {
				fmt.Println("error while insert VideoDetails into the Videos table")
				fmt.Println(err)
				
			}
			p := s.ErrorResult{Status: true, Message: "Video uploaded successfully", UserId: VideoDetails.UserId}
				json.NewEncoder(w).Encode(p)
		} else {
			p := s.ErrorResult{Status: true, Message: "video already exist"}
			json.NewEncoder(w).Encode(p)
			fmt.Println("video already exist")
		}

		// }
		// if e.ValidateEmail(VideoDetails.Mail) && VideoDetails.Etag != "" {
		// 	var etagfromDb string
		// 	resEtag := cassession.Session.Query("select etag from videos where etag=? allow filtering", VideoDetails.Etag)
		// 	resEtag.Scan(&etagfromDb)

		// 	if etagfromDb != VideoDetails.Etag {
		// 		res := cassession.Session.Query("select uid from signup where usermail=? allow filtering;", VideoDetails.Mail)
		// 		res.Scan(&VideoDetails.UserId)
		// 		fmt.Println(VideoDetails.UserId)

		// 		if err := cassession.Session.Query("insert into videos (videouid,videolink,videosizeinmb,title,description,language,genres,tags,agegroup,createddatetime,useruid,thumnail,etag)values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
		// 			VideoDetails.VideouID, VideoDetails.VideoLink, VideoDetails.VideoSizeInMb, VideoDetails.Title, VideoDetails.Description, VideoDetails.Language, VideoDetails.Genres, VideoDetails.Tags, VideoDetails.AgeGroup, VideoDetails.Createddatetime, VideoDetails.UserId, VideoDetails.Thumnail, VideoDetails.Etag).Exec(); err != nil {
		// 			fmt.Println("error while insert VideoDetails into the Videos table")
		// 			fmt.Println(err)
		// 		}
		// 	} else {
		// 		fmt.Println("video already exist")
		// 	}

		// } else if e.ValidateMobile(VideoDetails.Mail) && VideoDetails.Etag != "" {
		// 	var etagfromDb string
		// 	resEtag := cassession.Session.Query("select etag from videos where etag=? allow filtering", VideoDetails.Etag)
		// 	resEtag.Scan(&etagfromDb)
		// 	if etagfromDb != VideoDetails.Etag {
		// 		res := cassession.Session.Query("select uid from signup where mobile=? allow filtering;", VideoDetails.Mail)
		// 		res.Scan(&VideoDetails.UserId)
		// 		if err := cassession.Session.Query("insert into videos (videouid,videolink,videosizeinmb,title,description,language,genres,tags,agegroup,createddatetime,useruid,thumnail,etag)values(?,?,?,?,?,?,?,?,?,?,?,?,?)",
		// 			VideoDetails.VideouID, VideoDetails.VideoLink, VideoDetails.VideoSizeInMb, VideoDetails.Title, VideoDetails.Description, VideoDetails.Language, VideoDetails.Genres, VideoDetails.Tags, VideoDetails.AgeGroup, VideoDetails.Createddatetime, VideoDetails.UserId, VideoDetails.Thumnail, VideoDetails.Etag).Exec(); err != nil {
		// 			fmt.Println("error while insert VideoDetails into the Videos table")
		// 			fmt.Println(err)
		// 		}
		// 	} else {
		// 		fmt.Println("video already exist")
		// 	}
		// }

		// p := s.Result{Status: true, Message: "Video uploaded successfully", UserId: VideoDetails.UserId}
		// json.NewEncoder(w).Encode(p)
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
