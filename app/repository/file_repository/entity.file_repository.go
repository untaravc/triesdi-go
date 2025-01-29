package file_repository

type File struct {
	FileId           int    `json:"fileId"`
	FileUri          string `json:"fileUri"`
	FileThumbnailUri string `json:"fileThumbnailUri"`
}

type FileFilter struct {
	FileId  string   `json:"fileId"`
	FileIds []string `json:"fileIds"`
}
