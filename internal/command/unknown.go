package command

import (
	"tg-bot/internal/template"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandUnknown struct {
}

func NewCommandUnknown() *CommandUnknown {
	return &CommandUnknown{}
}

func (commandUnknown *CommandUnknown) Handle(update *telegram.Update, command string, args []string) (interface{}, error) {
	chatId, err := update.Message.GetChatId()
	if err != nil {
		return nil, err
	}

	text, err := template.NewUnknownTemplate(*update.Message.Text).GetText()
	if err != nil {
		return nil, err
	}

	return client.TelegramResponse(chatId, text)
}
