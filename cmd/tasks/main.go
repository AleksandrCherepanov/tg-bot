package main

import (
	"log"
	"net/http"

	config "github.com/AleksandrCherepanov/go_telegram/pkg/telegram/config"
	"github.com/AleksandrCherepanov/tg-bot/internal/middleware"
	"github.com/AleksandrCherepanov/tg-bot/internal/server"
	"github.com/gorilla/mux"
)

// TODO send telegram message concurrently
// TODO make config and client as a dependency of commands
// TODO get services as singletones
// TODO find out how to work with env file correctly
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
