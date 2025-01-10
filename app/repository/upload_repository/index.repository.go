package upload_repository

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3UploadRepository(accessKey, secretKey, region, bucket string) (*S3UploadRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3UploadRepository{
		s3Client: client,
		bucket:   bucket,
		region:   region,
	}, nil
}

func (r *S3UploadRepository) UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error) {
	_, err := r.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(fileName),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.bucket, r.region, fileName)
	return fileURL, nil
}
