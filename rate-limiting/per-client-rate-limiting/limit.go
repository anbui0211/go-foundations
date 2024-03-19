package clientratelimit

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)
type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time // Lần cuối cùng gọi API
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
			mu.Unlock()
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

