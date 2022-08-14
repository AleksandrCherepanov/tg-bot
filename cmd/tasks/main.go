package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Message struct {
	MessageID int  `json:"message_id"`
	Chat      Chat `json:"chat"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

func (c *Chat) getName() string {
	return c.FirstName + " " + c.LastName
}

func greetings(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(body))
	if len(body) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		return
	}

	update := &Update{}
	json.Unmarshal(body, update)

	request, err := http.NewRequest(
		"POST",
		"https://api.telegram.org/<token>/sendMessage",
		bytes.NewBuffer([]byte(`
{"chat_id": `+strconv.Itoa(update.Message.Chat.ID)+`, "text": "Hi, `+update.Message.Chat.getName()+`!"}`)),
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
	fmt.Println(string(rBody))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", greetings)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
