package telegram

type CallbackGame struct {
	UserId             int64   `json:"user_id"`
	Score              int64   `json:"score"`
	Force              *bool   `json:"force"`
	DisableEditMessage *bool   `json:"disable_edit_message"`
	ChatId             *int64  `json:"chat_id"`
	MessageId          *int64  `json:"message_id"`
	InlineMessageId    *string `json:"inline_message_id"`
}
