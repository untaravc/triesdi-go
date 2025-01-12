package upload_service

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	upload_repository "triesdi/app/repository/upload_repository"

	"github.com/google/uuid"
)

type UploadService struct {
	repo upload_repository.S3UploadRepository
}

func NewUploadService(repo upload_repository.S3UploadRepository) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) UploadFile(file multipart.File, header *multipart.FileHeader) (string, error) {
	uuid := uuid.New().String()
	extension := ""
	if dotIndex := strings.LastIndex(header.Filename, "."); dotIndex > 0 {
		extension = header.Filename[dotIndex:]
	}
	teamName := "tried_di"
	fileName := fmt.Sprintf("%s_%s_%d%s", teamName, uuid, time.Now().Unix(), extension)
	return s.repo.UploadFile(context.TODO(), file, fileName)
}
