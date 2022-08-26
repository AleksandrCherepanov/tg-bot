package command

import (
	"fmt"
	"tg-bot/internal/telegram"
)

type HandlerInterface interface {
	Handle(update *telegram.Update) (interface{}, error)
}

type CommandHandler struct {
	handlers map[string]HandlerInterface
}

func NewCommandHandler() *CommandHandler {
	handler := &CommandHandler{}
	handler.handlers = map[string]HandlerInterface{
		"/start": NewCommandStart(),
		"/help":  NewCommandHelp(),
	}

	return handler
}

func (commandHandler *CommandHandler) Handle(update *telegram.Update) (interface{}, error) {
	command := *update.Message.Text

	handler, ok := commandHandler.handlers[command]
	if !ok {
		return nil, fmt.Errorf("Unknown command: %v", command)
	}

	return handler.Handle(update)
}
