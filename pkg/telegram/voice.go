package telegram

type Voice struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	Duration     int64   `json:"duration"`
	MimeType     *string `json:"mime_type"`
	FileSize     *int64  `json:"file_size"`
}
