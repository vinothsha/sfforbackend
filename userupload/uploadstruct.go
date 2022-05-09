package userupload

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UploadVideo struct {
	VideoID   string `json:"videoid"`
	VideoLink string `json:"videolink"`
	// VideoLength   string `json:"videolength"`
	VideoSizeInMb float64 `json:"videosizeinmb"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Language      string  `json:"language"`
	Genres        string  `json:"geners"`
	AgeGroup      string  `json:"agegroup"`
	Mail          string  `json:"mail"`
	// Tags               []string `json:"tags"`
	Createddatetime string `json:"createddatetime"`
	// Lastupdatedatetime string `json:"lastupdatedatetime"`
}

const (
	BUCKET_NAME = "sfvideosplays"
	REGION      = "ap-south-1"

	// FILE      = "movie1.mp4"
	PART_SIZE = 6_000_000 // Has to be 5_000_000 minimim
	RETRIES   = 2
)

var (
	s3session *s3.S3
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}
