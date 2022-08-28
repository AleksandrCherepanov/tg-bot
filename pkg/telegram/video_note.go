package telegram

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int64      `json:"length"`
	Duration     int64      `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     *int64     `json:"file_size"`
}
