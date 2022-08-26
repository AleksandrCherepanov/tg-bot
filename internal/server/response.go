package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ResponseError(w http.ResponseWriter, data interface{}) {
	log.Printf("ERROR: %v\n", data)
	ResponseJson(w, "{}")
}
