package telegram

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     *string    `json:"file_name"`
	MimeType     *string    `json:"mime_type"`
	FileSize     *int64     `json:"file_size"`
}
