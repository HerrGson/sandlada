package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on port 7000")

	http.HandleFunc("/api/message", messageHandler)

	err := http.ListenAndServe(":7000", nil)
	if err != nil {
		log.Println(err)
	}
}

type Message struct {
	Content string `json:"content"`
}

func NewMessage(content string) Message {
	return Message{
		Content: content,
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s request was made\n", r.Method)
	msg := NewMessage("Hello123")
	bytes, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
