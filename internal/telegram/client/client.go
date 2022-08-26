package client

import (
	"bytes"
	"net/http"
	"strconv"
)

type TelegramHttpClient struct {
	client *http.Client
	host   string
}

func NewClient() *TelegramHttpClient {
	telegramHttpClient := &TelegramHttpClient{}
	telegramHttpClient.host = "https://api.telegram.org"
	telegramHttpClient.client = http.DefaultClient

	return telegramHttpClient
}

func (thc *TelegramHttpClient) SendMessage(userId int64, text string) (*http.Response, error) {
	request, err := http.NewRequest(
		"POST",
		thc.host+"/bot/sendMessage",
		bytes.NewBuffer([]byte(`{"chat_id": `+strconv.FormatInt(userId, 10)+`, "text": "`+text+`"}`)),
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
