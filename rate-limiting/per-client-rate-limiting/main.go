package clientratelimit

import (
	"encoding/json"
	"golang.org/x/time/rate"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time // Lần cuối cùng gọi API
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

// middleware để giới hạn tốc độ yêu cầu cho mỗi client dựa trên địa chỉ IP
func perClientRateLimiter(next func(writer http.ResponseWriter, request *http.Request)) http.Handler {
	var (
		mu      sync.Mutex
		clients = make(map[string]*Client)
	)

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get IP address from where the request is being made
		// If we have the IP address, we know how many times that particular IP address calls the API
		// and you want to limit that particular IP address and not all1 the other IP
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mu.Lock()
		if _, found := clients[ip]; !found {
			clients[ip] = &Client{limiter: rate.NewLimiter(2, 4)}
		}

		clients[ip].lastSeen = time.Now()
		if !clients[ip].limiter.Allow() {
			mu.Unlock()

			message := Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}
		mu.Unlock()
		// execute the next function
		next(w, r)
	})
}

func Main() {
	http.Handle("/ping", perClientRateLimiter(endpointHandler))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("There wa an error listening on port:8000", err)
	}
}
