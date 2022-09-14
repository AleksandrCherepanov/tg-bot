package client

import (
	"io"
	"tg-bot/pkg/config"
)

type TelegramResponse struct {
	ChatId  int64
	Text    string
	IsError bool
}

func NewTelegramResponse(chatId int64, text string, isError bool) TelegramResponse {
	return TelegramResponse{chatId, text, isError}
}

func (tr TelegramResponse) Error() string {
	if tr.IsError {
		return tr.Text
	}

	return ""
}

func (tr *TelegramResponse) Send() (interface{}, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	res, err := NewClient(cfg).SendMessage(tr.ChatId, tr.Text)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
