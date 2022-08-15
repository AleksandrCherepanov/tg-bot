package telegram

type File struct {
	fileId       string  `json:"file_id"`
	fileUniqueId string  `json:"file_unique_id"`
	fileSize     *int64  `json:"file_size"`
	filePath     *string `json:"file_path"`
}
