package command

import (
	"tg-bot/internal/template"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandHelp struct {
	chatId  int64
	message *telegram.Message
}

func NewCommandHelp(chatId int64, message *telegram.Message) *CommandHelp {
	return &CommandHelp{chatId, message}
}

func (c *CommandHelp) Handle(command string, args []string) (interface{}, error) {
	text, err := template.NewHelpTemplate().GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(c.chatId, text, false), nil
}
