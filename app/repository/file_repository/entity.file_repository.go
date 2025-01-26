package file_repository

type File struct {
	FileId           int    `json:"fileId"`
	FileUri          string `json:"fileUri"`
	FileThumbnailUri string `json:"fileThumbnailUri"`
}
