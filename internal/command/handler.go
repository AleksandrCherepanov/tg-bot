package command

import (
	"fmt"
	"strings"
	"tg-bot/pkg/telegram"
)

const unknownCommand = "/unknown"

type CommandHandlerInterface interface {
	Handle(update *telegram.Update, command string, args []string) (interface{}, error)
}

type CommandHandler struct {
	handlers map[string]CommandHandlerInterface
}

func NewCommandHandler() *CommandHandler {
	handler := &CommandHandler{}
	handler.handlers = map[string]CommandHandlerInterface{
		"/start":       NewCommandStart(),
		"/help":        NewCommandHelp(),
		"/l":           NewCommandList(),
		"/lc":          NewCommandList(),
		"/ls":          NewCommandList(),
		"/lg":          NewCommandList(),
		"/ld":          NewCommandList(),
		"/lda":         NewCommandList(),
		unknownCommand: NewCommandUnknown(),
	}

	return handler
}

func (commandHandler *CommandHandler) Handle(update *telegram.Update) (interface{}, error) {
	commandWithArgs := strings.Split(*update.Message.Text, " ")

	if len(commandWithArgs) == 0 {
		return nil, fmt.Errorf("Invalid command")
	}

	handler, ok := commandHandler.handlers[commandWithArgs[0]]
	if !ok {
		unknownHandler, _ := commandHandler.handlers[unknownCommand]
		return unknownHandler.Handle(update, unknownCommand, []string{})
	}

	return handler.Handle(update, commandWithArgs[0], commandWithArgs[1:])
}
