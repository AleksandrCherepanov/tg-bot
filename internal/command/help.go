package command

import (
	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram/client"
	"github.com/AleksandrCherepanov/tg-bot/internal/template"
)

type CommandHelp struct {
	chatId int64
}

func NewCommandHelp(chatId int64) *CommandHelp {
	return &CommandHelp{chatId}
}

func (c *CommandHelp) Handle(command string, args []string) (interface{}, error) {
	text, err := template.NewHelpTemplate().GetText()
	if err != nil {
		return nil, err
	}

	return client.NewTelegramResponse(c.chatId, text, false), nil
}
