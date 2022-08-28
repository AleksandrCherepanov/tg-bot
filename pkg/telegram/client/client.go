package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tg-bot/pkg/config"
)

type TelegramHttpClient struct {
	client *http.Client
	host   string
	config *config.Config
}

func NewClient(cfg *config.Config) *TelegramHttpClient {
	telegramHttpClient := &TelegramHttpClient{}
	telegramHttpClient.host = "https://api.telegram.org"
	telegramHttpClient.client = http.DefaultClient
	telegramHttpClient.config = cfg

	return telegramHttpClient
}

func (thc *TelegramHttpClient) SendMessage(userId int64, text string) (*http.Response, error) {
	message := NewMessageRequest(userId, text)

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsonMessage))
	request, err := http.NewRequest(
		"POST",
		thc.host+"/bot"+thc.config.Token+"/sendMessage",
		bytes.NewBuffer(jsonMessage),
	)
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
