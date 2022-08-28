package command

import (
	"tg-bot/pkg/telegram"
)

const unknownCommand = "/unknown"

type HandlerInterface interface {
	Handle(update *telegram.Update) (interface{}, error)
}

type CommandHandler struct {
	handlers map[string]HandlerInterface
}

func NewCommandHandler() *CommandHandler {
	handler := &CommandHandler{}
	handler.handlers = map[string]HandlerInterface{
		"/start":       NewCommandStart(),
		"/help":        NewCommandHelp(),
		unknownCommand: NewCommandUnknown(),
	}

	return handler
}

func (commandHandler *CommandHandler) Handle(update *telegram.Update) (interface{}, error) {
	command := *update.Message.Text

	handler, ok := commandHandler.handlers[command]
	if !ok {
		unknownHandler, _ := commandHandler.handlers[unknownCommand]
		return unknownHandler.Handle(update)
	}

	return handler.Handle(update)
}
