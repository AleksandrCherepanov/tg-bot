package client

import (
	"bytes"
	"encoding/json"
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

type messageRequest struct {
	ChatId    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func newMessageRequest(chatId int64, text string) *messageRequest {
	mr := &messageRequest{}
	mr.ChatId = chatId
	mr.Text = text
	mr.ParseMode = "MarkdownV2"

	return mr
}

func (thc *TelegramHttpClient) SendMessage(userId int64, text string) (*http.Response, error) {
	message := newMessageRequest(userId, text)

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

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
