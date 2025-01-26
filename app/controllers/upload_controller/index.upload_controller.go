package upload_controller

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"triesdi/app/repository/file_repository"
	"triesdi/app/services/upload_service"

	"github.com/gin-gonic/gin"
)

const (
	MaxFileSize  = 100 * 1024 // 100 KiB
	MaxGoroutine = 5          // Max concurrent uploads
)

var supportedFormats = map[string]bool{
	".jpeg": true,
	".jpg":  true,
	".png":  true,
}

// UploadHandler handles file upload requests.
func UploadHandler(s3Uploader *upload_service.S3Uploader) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse multipart form
		if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
			return
		}

		form := c.Request.MultipartForm
		files := form.File["file"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
			return
		}

		// Goroutine control
		limitCh := make(chan struct{}, MaxGoroutine)
		wg := sync.WaitGroup{}
		errors := make(chan error, len(files))

		for _, fileHeader := range files {
			limitCh <- struct{}{}
			wg.Add(1)

			go func(fileHeader *multipart.FileHeader) {
				defer func() {
					<-limitCh
					wg.Done()
				}()

				// Validate file size
				if fileHeader.Size > MaxFileSize {
					errors <- fmt.Errorf("file %s exceeds size limit", fileHeader.Filename)
					return
				}

				// Validate file type
				ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
				if !supportedFormats[ext] {
					errors <- fmt.Errorf("file %s has unsupported format", fileHeader.Filename)
					return
				}

				// Open file
				file, err := fileHeader.Open()
				if err != nil {
					errors <- fmt.Errorf("failed to open file %s: %v", fileHeader.Filename, err)
					return
				}
				defer file.Close()

				// Read file content
				fileBytes := make([]byte, fileHeader.Size)
				_, err = file.Read(fileBytes)
				if err != nil {
					errors <- fmt.Errorf("failed to read file %s: %v", fileHeader.Filename, err)
					return
				}

				// Upload to S3
				err = s3Uploader.UploadFile(context.TODO(), fileBytes, fileHeader.Filename)
				if err != nil {
					errors <- fmt.Errorf("failed to upload file %s: %v", fileHeader.Filename, err)
					return
				}
			}(fileHeader)
		}

		wg.Wait()
		close(errors)

		// Collect errors
		var errMsgs []string
		for err := range errors {
			errMsgs = append(errMsgs, err.Error())
		}

		if len(errMsgs) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": errMsgs})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "All files uploaded successfully"})
	}
}

func AddFile(c *gin.Context) {
	// Parse multipart form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form"})
		return
	}

	form := c.Request.MultipartForm
	files := form.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	for _, fileHeader := range files {
		if fileHeader.Size > MaxFileSize {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit file size"})
			return
		}

		// Validate file type
		ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
		if !supportedFormats[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file has unsupported format"})
			return
		}
		result, err := file_repository.CreateDummy(ext)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "All files uploaded successfully"})
}
