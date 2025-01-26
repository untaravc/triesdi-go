package file_repository

import (
	"fmt"
	"triesdi/app/configs/aws_config"
	"triesdi/app/utils/common"
	"triesdi/app/utils/database"
)

const SAMPLE = "https://projectsprint-bucket-public-read.s3.ap-southeast-1.amazonaws.com/uploads/my-image.jpg"
const DB_NAME = "files"

func CreateDummy(file_type string) (File, error) {
	filePath := "https://"
	filePath += aws_config.S3_BUCKET
	filePath += ".s3."
	filePath += aws_config.AWS_REGION
	filePath += ".amazonaws.com/uploads/"

	fileUri := filePath + common.RandomString(60) + file_type
	fileThumbnailUri := filePath + "thumbnail/" + common.RandomString(60) + file_type

	file := File{FileUri: fileUri, FileThumbnailUri: fileThumbnailUri}

	query := fmt.Sprintf("INSERT INTO %s (file_uri, file_thumbnail_uri) VALUES ($1, $2) RETURNING file_id", DB_NAME)

	var insertedID int
	err := database.DB.QueryRow(query, file.FileUri, file.FileThumbnailUri).Scan(&insertedID)
	if err != nil {
		return file, err
	}

	file.FileId = insertedID

	return file, nil
}

func GetById(id string) (File, error) {
	file := File{}

	query := fmt.Sprintf("SELECT file_id, file_uri, file_thumbnail_uri FROM %s WHERE file_id = $1", DB_NAME)

	err := database.DB.QueryRow(query, id).Scan(&file.FileId, &file.FileUri, &file.FileThumbnailUri)
	if err != nil {
		return file, err
	}

	return file, nil
}
