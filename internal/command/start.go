package command

import (
	"tg-bot/internal/template"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandStart struct {
}

func NewCommandStart() *CommandStart {
	return &CommandStart{}
}

func (commandStart *CommandStart) Handle(update *telegram.Update, command string, args []string) (interface{}, error) {
	chatId, err := update.Message.GetChatId()
	if err != nil {
		return nil, err
	}

	text, err := template.NewStartTemplate().GetText()
	if err != nil {
		return nil, err
	}

	return client.TelegramResponse(chatId, text)
}
