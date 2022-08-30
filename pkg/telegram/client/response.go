package client

import (
	"fmt"
	"io"
	"tg-bot/pkg/config"
)

func TelegramResponse(chatId int64, data interface{}) (interface{}, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	text, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("Can't parse telegram data")
	}

	res, err := NewClient(cfg).SendMessage(chatId, text)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
