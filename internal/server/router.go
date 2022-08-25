package server

import (
	"encoding/json"
	"io"
	"net/http"
	"tg-bot/internal/command"
	"tg-bot/internal/telegram"
)

type Router struct {
	handlers map[string]command.HandlerInterface
}

func NewRouter() *Router {
	router := &Router{}
	router.handlers = map[string]command.HandlerInterface{
		"command": command.NewCommandHandler(),
	}

	return router
}

func (router *Router) Resolve(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	update := &telegram.Update{}
	err = json.Unmarshal(body, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := update.Message
	if message == nil {
		http.Error(w, "Can't process message", http.StatusInternalServerError)
		return
	}

	if message.Entities == nil {
		http.Error(w, "Can't process message", http.StatusInternalServerError)
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
		http.Error(w, handleError.Error(), http.StatusInternalServerError)
		return
	}

	ResponseJson(w, result)
}
