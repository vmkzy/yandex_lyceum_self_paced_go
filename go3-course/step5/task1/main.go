package main

import (
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("incoming request", slog.String("method", r.Method), slog.String("path", r.URL.Path))
		next.ServeHTTP(w, r)
	})
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, middleware!"))
}
func main() {
	mux := http.NewServeMux()
	hello := http.HandlerFunc(helloHandler)
	mux.Handle("/hello", Logger(hello))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Server fail")
	}
}
