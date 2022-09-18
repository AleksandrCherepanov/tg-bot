package server

import (
	"encoding/json"
	"net/http"

	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram"
	"github.com/AleksandrCherepanov/go_telegram/pkg/telegram/client"
	"github.com/AleksandrCherepanov/tg-bot/internal/command"
)

type Router struct {
	handlers map[string]command.HandlerInterface
}

func NewRouter() *Router {
	return &Router{}
}

func (router *Router) WithHandlers(chatId int64, message *telegram.Message) *Router {
	router.handlers = map[string]command.HandlerInterface{
		"command": command.NewCommandHandler(chatId, message),
	}
	return router
}

func (router *Router) Resolve(w http.ResponseWriter, req *http.Request) {
	body, ok := GetParsedBody(req)
	if !ok {
		ResponseError(w, "Can't get parsed body")
		return
	}

	update := &telegram.Update{}
	err := json.Unmarshal(body, update)
	if err != nil {
		ResponseError(w, err.Error())
		return
	}

	message := update.Message
	if message == nil {
		ResponseError(w, "Can't process message")
		return
	}

	if message.Entities == nil {
		ResponseError(w, "Can't process message")
		return
	}

	chatId, err := update.Message.GetChatId()
	if err != nil {
		ResponseError(w, "Can't process message")
	}

	var result interface{}
	var handleError error
	for _, entity := range *&message.Entities {
		if entity.IsCommand() {
			router = router.WithHandlers(chatId, message)
			result, handleError = router.handlers["command"].Handle(chatId, update.Message)
		}
	}

	if handleError != nil {
		tgResponse, ok := handleError.(client.TelegramResponse)
		if ok {
			_, err = tgResponse.Send()
			if err != nil {
				ResponseError(w, err.Error())
				return
			}
		}
		ResponseError(w, handleError.Error())
		return
	}

	if tgResult, ok := result.(client.TelegramResponse); ok {
		_, err = tgResult.Send()
		if err != nil {
			ResponseError(w, err.Error())
			return
		}
	}
	ResponseJson(w, result)
}
