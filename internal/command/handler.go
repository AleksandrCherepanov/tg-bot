package command

import (
	"strings"

	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram"
	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram/client"
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
		"/help":        NewCommandHelp(chatId),
		"/l":           NewCommandList(chatId),
		"/lc":          NewCommandList(chatId),
		"/ls":          NewCommandList(chatId),
		"/lg":          NewCommandList(chatId),
		"/ld":          NewCommandList(chatId),
		"/lda":         NewCommandList(chatId),
		"/t":           NewCommandTask(chatId),
		"/tc":          NewCommandTask(chatId),
		"/td":          NewCommandTask(chatId),
		"/tda":         NewCommandTask(chatId),
		"/tm":          NewCommandTask(chatId),
		"/tma":         NewCommandTask(chatId),
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
