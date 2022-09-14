package command

import (
	"strings"
	"tg-bot/pkg/telegram"
	"tg-bot/pkg/telegram/client"
)

const unknownCommand = "/unknown"

type HandlerInterface interface {
	Handle(chatId int64, message *telegram.Message) (interface{}, error)
}

type CommandHandlerInterface interface {
	Handle(command string, args []string) (interface{}, error)
}

type CommandHandler struct {
	handlers map[string]CommandHandlerInterface
}

func NewCommandHandler(chatId int64, message *telegram.Message) *CommandHandler {
	commandHandler := &CommandHandler{}
	commandHandler.handlers = map[string]CommandHandlerInterface{
		"/start":       NewCommandStart(chatId, message),
		"/help":        NewCommandHelp(chatId, message),
		"/l":           NewCommandList(chatId, message),
		"/lc":          NewCommandList(chatId, message),
		"/ls":          NewCommandList(chatId, message),
		"/lg":          NewCommandList(chatId, message),
		"/ld":          NewCommandList(chatId, message),
		"/lda":         NewCommandList(chatId, message),
		unknownCommand: NewCommandUnknown(chatId, message),
	}

	return commandHandler
}

func (commandHandler *CommandHandler) Handle(chatId int64, message *telegram.Message) (interface{}, error) {
	commandWithArgs := strings.Split(*message.Text, " ")

	if len(commandWithArgs) == 0 {
		return nil, client.NewTelegramResponse(chatId, "Invalid command", true)
	}

	handler, ok := commandHandler.handlers[commandWithArgs[0]]
	if !ok {
		unknownHandler, _ := commandHandler.handlers[unknownCommand]
		return unknownHandler.Handle(unknownCommand, []string{})
	}

	return handler.Handle(commandWithArgs[0], commandWithArgs[1:])
}
