package server

import (
	"encoding/json"
	"net/http"
	"tg-bot/internal/command"
	"tg-bot/pkg/telegram"
)

type Router struct {
	handlers map[string]HandlerInterface
}

type HandlerInterface interface {
	Handle(update *telegram.Update) (interface{}, error)
}

func NewRouter() *Router {
	router := &Router{}
	router.handlers = map[string]HandlerInterface{
		"command": command.NewCommandHandler(),
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

	var result interface{}
	var handleError error
	for _, entity := range *&message.Entities {
		if entity.IsCommand() {
			result, handleError = router.handlers["command"].Handle(update)
		}
	}

	if handleError != nil {
		ResponseError(w, handleError.Error())
		return
	}

	ResponseJson(w, result)
}
