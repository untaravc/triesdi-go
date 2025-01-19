package upload_service

import (
	"context"
	"mime/multipart"

	upload_repository "triesdi/app/repository/upload_repository"
)

type UploadService struct {
	repo upload_repository.S3UploadRepository
}

func NewUploadService(repo upload_repository.S3UploadRepository) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) UploadFile(file multipart.File, header *multipart.FileHeader, fileName string) (string, error) {
	return s.repo.UploadFile(context.TODO(), file, fileName)
}
