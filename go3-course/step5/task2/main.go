package main

import (
	"context"
	"net/http"
	"regexp"
)

type contextKey string

const nameKey contextKey = "userName"

var check = regexp.MustCompile("^[a-zA-Z]+$")

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		ctx := context.WithValue(r.Context(), nameKey, name)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, ok := r.Context().Value(nameKey).(string)
		if ok && name != "stranger" {
			if !check.MatchString(name) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("hello dirty hacker"))
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Context().Value(nameKey).(string)
	w.Write([]byte("hello " + name))
}

func main() {
	mux := http.NewServeMux()
	hello := SetDefaultName(Sanitize(HelloHandler))

	mux.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", mux)
}
