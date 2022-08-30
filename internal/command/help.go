package command

import (
	"tg-bot/internal/template"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandHelp struct {
}

func NewCommandHelp() *CommandHelp {
	return &CommandHelp{}
}

func (commandHelp *CommandHelp) Handle(update *telegram.Update, command string, args []string) (interface{}, error) {
	chatId, err := update.Message.GetChatId()
	if err != nil {
		return nil, err
	}

	text, err := template.NewHelpTemplate().GetText()
	if err != nil {
		return nil, err
	}

	return client.TelegramResponse(chatId, text)
}
