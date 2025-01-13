package upload_requst

import (
	"mime/multipart"
)

type UploadRequest struct {
	File *multipart.File `validate:"required"`
}
