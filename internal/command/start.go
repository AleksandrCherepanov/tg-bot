package command

import (
	"tg-bot/internal/template"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandStart struct {
	chatId  int64
	message *telegram.Message
}

func NewCommandStart(chatId int64, message *telegram.Message) *CommandStart {
	return &CommandStart{chatId, message}
}

func (c *CommandStart) Handle(command string, args []string) (interface{}, error) {
	text, err := template.NewStartTemplate().GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(c.chatId, text, false), nil
}
