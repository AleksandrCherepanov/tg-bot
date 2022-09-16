package command

import (
	"tg-bot/internal/template"
	"tg-bot/internal/user"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

type CommandStart struct {
	chatId      int64
	message     *telegram.Message
	userStorage *user.UserStorage
}

func NewCommandStart(chatId int64, message *telegram.Message) *CommandStart {
	return &CommandStart{
		chatId:      chatId,
		message:     message,
		userStorage: user.GetUserStorage(),
	}
}

func (c *CommandStart) Handle(command string, args []string) (interface{}, error) {
	text, err := template.NewStartTemplate().GetText()
	if err != nil {
		return nil, err
	}

	c.userStorage.CreateUser(c.chatId, c.message.Chat.GetName())
	return client.NewTelegramResponse(c.chatId, text, false), nil
}
