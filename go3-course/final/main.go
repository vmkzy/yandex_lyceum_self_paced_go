package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func NewStore() *Store {
	return &Store{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *Store) CreateUser(name string, age int) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:   s.nextID,
		Name: name,
		Age:  age,
	}
	s.users[user.ID] = user
	s.nextID++

	return user
}

func (s *Store) GetUser(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[id]
	return user, ok
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		logger.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.statusCode),
			slog.Duration("duration", time.Since(start)))
	})
}

type createUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func writeBadRequest(w http.ResponseWriter) {
	http.Error(w, "bad request", http.StatusBadRequest)
}

func writeNotFound(w http.ResponseWriter) {
	http.Error(w, "not found", http.StatusNotFound)
}

func createUserHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req createUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeBadRequest(w)
			return
		}

		user := store.CreateUser(req.Name, req.Age)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	}
}

func getUserHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		idPart := strings.TrimPrefix(r.URL.Path, "/users/")
		if idPart == "" || strings.Contains(idPart, "/") {
			writeNotFound(w)
			return
		}

		id, err := strconv.Atoi(idPart)
		if err != nil || id < 1 {
			writeBadRequest(w)
			return
		}

		user, ok := store.GetUser(id)
		if !ok {
			writeNotFound(w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	}
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	store := NewStore()

	mux := http.NewServeMux()
	mux.HandleFunc("/users", createUserHandler(store))
	mux.HandleFunc("/users/", getUserHandler(store))

	handler := loggingMiddleware(logger, mux)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
