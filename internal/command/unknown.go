package command

import (
	"io"
	"tg-bot/internal/template"
	"tg-bot/pkg/config"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandUnknown struct {
}

func NewCommandUnknown() *CommandUnknown {
	return &CommandUnknown{}
}

func (commandUnknown *CommandUnknown) Handle(update *telegram.Update) (interface{}, error) {
	chatId, err := update.Message.GetChatId()
	if err != nil {
		return nil, err
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	text, err := template.NewUnknownTemplate(*update.Message.Text).GetText()
	if err != nil {
		return nil, err
	}

	res, err := client.NewClient(cfg).SendMessage(chatId, text)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
