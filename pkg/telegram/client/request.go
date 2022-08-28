package client

type MessageRequest struct {
	ChatId    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func NewMessageRequest(chatId int64, text string) *MessageRequest {
	mr := &MessageRequest{}
	mr.ChatId = chatId
	mr.Text = text
	mr.ParseMode = "MarkdownV2"

	return mr
}
