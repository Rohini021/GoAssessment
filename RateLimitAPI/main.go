package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu     sync.Mutex
	counts map[string]int
	reset  map[string]time.Time
	limit  int
}

func rateLimiterMethod(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		counts: make(map[string]int),
		reset:  make(map[string]time.Time),
		limit:  limit,
	}
}

func (rl *RateLimiter) checkCounter(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	if lastReset, exists := rl.reset[ip]; !exists || now.After(lastReset.Add(time.Minute)) {
		rl.counts[ip] = 0
		rl.reset[ip] = now
	}
	rl.counts[ip]++
	return rl.counts[ip] <= rl.limit
}

var limiter = rateLimiterMethod(3, time.Minute)

func processRequest(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	if !limiter.checkCounter(ip) {
		fmt.Println("exceeded rate limit")
		http.Error(w, "exceeded rate limit", http.StatusTooManyRequests)
		return
	}
	fmt.Println("request success")
	fmt.Fprint(w, "successfully request received")
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			return ip
		}
		return ""
	}
	return ip
}

func main() {
	http.HandleFunc("/api/rate/post", processRequest)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// To test above code hit command on terminal after running go code : curl -X POST http://localhost:8080/api/rate/post
// after hitting API for 3 times, on 4th time it will show error as limit exceeded as we have set limit of 3 per minute
