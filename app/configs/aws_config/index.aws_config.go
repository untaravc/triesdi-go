package aws_config

import (
	"os"
)

var AWS_ACCESS_KEY_ID = ""
var AWS_SECRET_ACCESS_KEY = ""
var AWS_REGION = ""
var S3_BUCKET = ""

func InitAwsConfig() {
	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")

	if awsAccessKeyId != "" {
		AWS_ACCESS_KEY_ID = awsAccessKeyId
	}

	if awsSecretAccessKey != "" {
		AWS_SECRET_ACCESS_KEY = awsSecretAccessKey
	}

	if awsRegion != "" {
		AWS_REGION = awsRegion
	}

	if s3Bucket != "" {
		S3_BUCKET = s3Bucket
	}
}
