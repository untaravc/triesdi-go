package v1_upload_controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

// UploadImage handles the POST /v1/upload route, where the client will upload a file to
// the server. The server will then upload the file to the AWS S3 bucket provided in the
// environment variables. The server will return a JSON response with the file URL, which
// can be used to access the uploaded file.
func UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file: " + err.Error()})
		return
	}
	defer file.Close()

	uuid := uuid.New().String()
	extension := ""
	if dotIndex := strings.LastIndex(header.Filename, "."); dotIndex > 0 {
		extension = header.Filename[dotIndex:]
	}
	team_name := "tried_di"
	fileName := fmt.Sprintf("%s_%s_%d%s", team_name, uuid, time.Now().Unix(), extension)

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")

	if awsAccessKey == "" || awsSecretKey == "" || awsRegion == "" || s3Bucket == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Missing AWS credentials"})
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, ""),
		),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load AWS config: " + err.Error()})
		return
	}

	s3Client := s3.NewFromConfig(cfg)

	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s3Bucket, awsRegion, fileName)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully!",
		"fileURL": fileURL,
	})
}
