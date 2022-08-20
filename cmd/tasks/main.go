package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"tg-bot/internal/task"
	"tg-bot/internal/telegram"
)

func greetings(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("REQUEST:")
	fmt.Println(string(body))
	fmt.Println()
	if len(body) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	update := &telegram.Update{}
	json.Unmarshal(body, update)

	chatId := update.Message.Chat.Id
	name := update.Message.Chat.GetName()

	request, err := http.NewRequest(
		"POST",
		"https://api.telegram.org/bot/sendMessage",
		bytes.NewBuffer([]byte(`{"chat_id": `+strconv.FormatInt(chatId, 10)+`, "text": "Hi, `+name+`!"}`)),
	)
	request.Header.Add("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rBody, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("RESPONSE:")
	fmt.Println(string(rBody))
	fmt.Println()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	server := task.NewTaskServer()
	mux := http.NewServeMux()
	//mux.HandleFunc("/tasks", greetings)
	mux.HandleFunc("/task/", server.Handler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
