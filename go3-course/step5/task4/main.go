package main

import (
	"fmt"
	"net/http"
	"strings"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}
func answerHandler(w http.ResponseWriter, r *http.Request) {
	username, _, _ := r.BasicAuth()
	fmt.Fprintf(w, "Welcome, %s!", username)
}
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/answer/", Authorization(answerHandler))
	http.ListenAndServe(":8080", mux)
}
