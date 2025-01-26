package upload_service

import (
	"bytes"
	"context"
	"errors"
	"triesdi/app/configs/aws_config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Uploader struct handles file uploads to S3.
type S3Uploader struct {
	client *s3.Client
}

// NewS3Uploader initializes and returns an S3Uploader instance.
func NewS3Uploader(client *s3.Client) *S3Uploader {
	return &S3Uploader{client: client}
}

// UploadFile uploads the file to S3.
func (u *S3Uploader) UploadFile(ctx context.Context, fileBytes []byte, fileName string) error {
	if len(fileBytes) == 0 || fileName == "" {
		return errors.New("invalid file input")
	}

	uploader := manager.NewUploader(u.client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(aws_config.S3_BUCKET),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileBytes),
	})
	return err
}
