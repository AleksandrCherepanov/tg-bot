package client

import (
	"encoding/json"
	"fmt"
	"io"
	"tg-bot/pkg/config"
)

type TelegramResponse struct {
	ChatId       int64
	Text         string
	IsError      bool
	IsPinMessage bool
}

func NewTelegramResponse(chatId int64, text string, isError bool) TelegramResponse {
	return TelegramResponse{chatId, text, isError, false}
}

func (tr *TelegramResponse) WithPinMessage() *TelegramResponse {
	tr.IsPinMessage = true
	return tr
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

	tgClient := NewClient(cfg)
	res, err := tgClient.SendMessage(tr.ChatId, tr.Text)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if tr.IsPinMessage {
		responseMessage := &MessageResponse{}
		err := json.Unmarshal(responseBody, responseMessage)
		if err != nil {
			return nil, err
		}

		if !responseMessage.Ok {
			return nil, fmt.Errorf("Something went wrong")
		}

		_, err = tgClient.unpinAllChatMessages(tr.ChatId)
		if err != nil {
			return nil, err
		}

		_, err = tgClient.PinMessage(tr.ChatId, responseMessage.Result.MessageId, true)
		if err != nil {
			return nil, err
		}
	}

	return responseBody, nil
}
