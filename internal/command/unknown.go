package command

import (
	"tg-bot/internal/template"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandUnknown struct {
	chatId  int64
	message *telegram.Message
}

func NewCommandUnknown(chatId int64, message *telegram.Message) *CommandUnknown {
	return &CommandUnknown{chatId, message}
}

func (c *CommandUnknown) Handle(command string, args []string) (interface{}, error) {
	text, err := template.NewUnknownTemplate(*c.message.Text).GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(c.chatId, text, true), nil
}
