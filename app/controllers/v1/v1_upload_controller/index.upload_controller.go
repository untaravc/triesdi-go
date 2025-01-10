package v1_upload_controller

import (
	"net/http"
	"os"
	"triesdi/app/responses/response"

	upload_repository "triesdi/app/repository/upload_repository"
	upload "triesdi/app/service/upload"

	"github.com/gin-gonic/gin"
)

// UploadImage handles the POST /v1/upload route, where the client will upload a file to
// the server. The server will then upload the file to the AWS S3 bucket provided in the
// environment variables. The server will return a JSON response with the file URL, which
// can be used to access the uploaded file.

func UploadImage(ctx *gin.Context) {

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")

	uploadRepo, err := upload_repository.NewS3UploadRepository(awsAccessKey, awsSecretKey, awsRegion, s3Bucket)
	if err != nil {
		panic("Failed to initialize upload repository: " + err.Error())
	}

	file, header, err := ctx.Request.FormFile("file")
	uploadService := upload.NewUploadService(*uploadRepo)

	if err != nil {
		response.BaseResponse(ctx, http.StatusBadRequest, false, "Please insert file to upload", "Failed to get file: "+err.Error())
		return
	}
	defer file.Close()

	fileURL, err := uploadService.UploadFile(file, header)
	if err != nil {
		response.BaseResponse(ctx, http.StatusInternalServerError, false, "Internal System Error", "Failed to get file: "+err.Error())
		return
	}

	response.BaseResponse(ctx, http.StatusOK, true, "File uploaded successfully!", fileURL)

}
