package command

import (
	"io"
	"tg-bot/internal/telegram"
	"tg-bot/internal/telegram/client"
)

type CommandStart struct {
}

func NewCommandStart() *CommandStart {
	return &CommandStart{}
}

func (commandStart *CommandStart) Handle(update *telegram.Update) (interface{}, error) {
	chatId, err := update.Message.GetChatId()
	if err != nil {
		return nil, err
	}
	res, err := client.NewClient().SendMessage(chatId, "START COMMAND IS CALLED")
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
