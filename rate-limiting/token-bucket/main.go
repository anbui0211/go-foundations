package ratelimittokenbuket

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	message := &Message{
		Status: "Successfully",
		Body:   "HI! You've reached the API. How may I help you?",
	}

	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func Main() {
	http.Handle("/ping", rateLimiter(endpointHandler))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("There wa an error listening on port:8000", err)
	}
}

