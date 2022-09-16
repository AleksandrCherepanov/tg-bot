package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tg-bot/pkg/config"
	"tg-bot/pkg/telegram"
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

type MessageResponse struct {
	Ok     bool             `json:"ok"`
	Result telegram.Message `json:"result"`
}

type pinMessageRequest struct {
	ChatId              int64 `json:"chat_id"`
	MessageId           int64 `json:"message_id"`
	DisableNotificaiton bool  `json:"disable_notificaiton"`
}

type unpinAllChatMessages struct {
	ChatId int64 `json:"chat_id"`
}

func newMessageRequest(chatId int64, text string) *messageRequest {
	mr := &messageRequest{}
	mr.ChatId = chatId
	mr.Text = text
	mr.ParseMode = "MarkdownV2"

	return mr
}

func newPinMessageRequest(chatId int64, messageId int64, disableNotification bool) *pinMessageRequest {
	return &pinMessageRequest{
		ChatId:              chatId,
		MessageId:           messageId,
		DisableNotificaiton: disableNotification,
	}
}

func newUnpinAllChatMessages(chatId int64) *unpinAllChatMessages {
	return &unpinAllChatMessages{chatId}
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

func (thc *TelegramHttpClient) PinMessage(
	chatId int64,
	messageId int64,
	disableNotification bool,
) (*http.Response, error) {
	pinMessage := newPinMessageRequest(chatId, messageId, disableNotification)

	jsonMessage, err := json.Marshal(pinMessage)
	fmt.Println(string(jsonMessage))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		"POST",
		thc.host+"/bot"+thc.config.Token+"/pinChatMessage",
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

func (thc *TelegramHttpClient) unpinAllChatMessages(chatId int64) (*http.Response, error) {
	unpiAllChatMessages := newUnpinAllChatMessages(chatId)

	jsonMessage, err := json.Marshal(unpiAllChatMessages)
	fmt.Println(string(jsonMessage))
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(
		"POST",
		thc.host+"/bot"+thc.config.Token+"/pinChatMessage",
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
