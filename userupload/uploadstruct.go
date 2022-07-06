package userupload

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	BUCKET_NAME = "storyflicsvideos"
	REGION      = "ap-south-1"

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
