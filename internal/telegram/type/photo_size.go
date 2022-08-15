package telegram

type PhotoSize struct {
	fileId       string `json:"file_id"`
	fileUniqueId string `json:"file_unique_id"`
	width        int64  `json:"width"`
	height       int64  `json:"height"`
	fileSize     *int64 `json:"file_size"`
}
