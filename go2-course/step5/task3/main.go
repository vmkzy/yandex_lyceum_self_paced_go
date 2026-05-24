package main

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func Request(ctx context.Context, url string) (string, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, err
	}
	return string(body), resp.StatusCode, nil

}
func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	ans := make([]*APIResponse, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(indx int, URL string) {
			defer wg.Done()
			res := &APIResponse{URL: URL}
			data, statusCode, err := Request(ctx, URL)
			if err != nil {
				res.Err = err
			} else {
				res.Data = data
				res.StatusCode = statusCode
				res.Err = nil
			}
			ans[indx] = res
		}(i, url)
	}
	wg.Wait()
	return ans
}
