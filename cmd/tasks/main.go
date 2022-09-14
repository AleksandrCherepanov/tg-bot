package main

import (
	"log"
	"net/http"
	"tg-bot/internal/middleware"
	"tg-bot/internal/server"
	"tg-bot/pkg/config"

	"github.com/gorilla/mux"
)

// TODO fix without space list name
// TODO create user only for start command
// TODO make commands like `command`
// TODO return unknown commands to telegram as an answer
// TODO make config and client as a dependency of commands
// TODO get services as singletones
// TODO find out how to work with env file correctly
// TODO remove debug logs
func main() {
	_, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Can't intitalize config. %v", err.Error())
	}

	server := server.NewRouter()
	router := mux.NewRouter()

	router.HandleFunc("/tasks", server.Resolve).Methods("POST", "GET")

	loggedRouter := middleware.Logging(router)
	panicRecoveryRouter := middleware.PanicRecovery(loggedRouter)
	log.Fatal(http.ListenAndServe(":3000", panicRecoveryRouter))
}
