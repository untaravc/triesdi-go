package v1_upload_controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"triesdi/app/responses/response"

	upload_repository "triesdi/app/repository/upload_repository"
	upload "triesdi/app/service/upload"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UploadImage handles the POST /v1/upload route, where the client will upload a file to
// the server. The server will then upload the file to the AWS S3 bucket provided in the
// environment variables. The server will return a JSON response with the file URL, which
// can be used to access the uploaded file.

var uploadChannel = make(chan struct{}, 10)

func UploadImage(ctx *gin.Context) {

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")

	uploadRepo, err := upload_repository.UploadRepository(awsAccessKey, awsSecretKey, awsRegion, s3Bucket)
	if err != nil {
		panic("Failed to initialize upload repository: " + err.Error())
	}

	file, header, err := ctx.Request.FormFile("file")
	uploadService := upload.NewUploadService(*uploadRepo)

    if err != nil {
        response.BaseResponse(ctx, http.StatusBadRequest, false, "Please insert file to upload", gin.H{"error" :err.Error()})
        return
    }
    defer file.Close()

    if header.Size > 100*1024 {
        response.BaseResponse(ctx, http.StatusBadRequest, false, "File too large", gin.H{"error" : "File size should not exceed 100KB"})
        return
    }

    fileExtension := strings.ToLower(filepath.Ext(header.Filename))
    if fileExtension != ".jpg" && fileExtension != ".png" && fileExtension != ".jpeg" {
        response.BaseResponse(ctx, http.StatusBadRequest, false, "Invalid file format", gin.H{"error": "Only JPG & PNG Allowed"})
        return
    }

	// CREATE URI
	uuid := uuid.New().String()
	extension := ""
	if dotIndex := strings.LastIndex(header.Filename, "."); dotIndex > 0 {
		extension = header.Filename[dotIndex:]
	}
	teamName := "tried_di"
	fileName := fmt.Sprintf("%s_%s_%d%s", teamName, uuid, time.Now().Unix(), extension)

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s3Bucket, awsRegion, fileName)

    uploadChannel <- struct{}{} // Acquire a slot
    go func() {
        defer func() { <-uploadChannel }() // Release the slot

        _, err := uploadService.UploadFile(file, header, fileName)
        if err != nil {
            response.BaseResponse(ctx, http.StatusInternalServerError, false, "Internal System Error", gin.H{"error": err.Error()})
            return
        }
    }()
	
	response.BaseResponse(ctx, http.StatusOK, true, "OK", gin.H{"uri": fileURL})
}


