package main

import (
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	readHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "http://localhost:8081/provideData", nil)
		if err != nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}

		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
			return
		}

		_, _ = w.Write(data)
	})
	timeoutHandler := http.TimeoutHandler(
		readHandler, maxTimeout, "Service unavailable",
	)

	mux := http.NewServeMux()
	mux.Handle("/readSource", timeoutHandler)
	_ = http.ListenAndServe(":8080", mux)
}
