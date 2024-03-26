package tollboothratelimit

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/didip/tollbooth/v7"
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
	message := Message{
		Status: "Request Failed",
		Body:   "The API is at capacity, try again later.",
	}

	jsonMessage, _ := json.Marshal(message)

	tlbthLimiter := tollbooth.NewLimiter(1, nil) // giới hạn là 1 request mỗi giây
	tlbthLimiter.SetMessageContentType("application/json")
	tlbthLimiter.SetMessage(string(jsonMessage))

	http.Handle("/ping", tollbooth.LimitFuncHandler(tlbthLimiter, endpointHandler))

	// start server
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("There wa an error listening on port:8000", err)
	}
}
