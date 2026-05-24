package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Fib struct {
	mu    sync.Mutex
	prev  int
	next  int
	count int
}

func NewFib() *Fib {
	return &Fib{
		prev: 0,
		next: 1,
	}
}

func (s *Fib) FibHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Fprint(w, s.prev)
	s.prev, s.next = s.next, s.prev+s.next
}
func (s *Fib) MetricsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.count++
	})
}

func main() {
	server := NewFib()
	mux := http.NewServeMux()
	fib := http.HandlerFunc(server.FibHandler)
	mux.HandleFunc("/", fib)
	mux.Handle("/metrics", server.MetricsHandler(fib))
	http.ListenAndServe(":8080", mux)
}
