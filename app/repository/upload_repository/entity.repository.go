package upload_repository

import "github.com/aws/aws-sdk-go-v2/service/s3"

type S3UploadRepository struct {
	s3Client *s3.Client
	bucket   string
	region   string
}
