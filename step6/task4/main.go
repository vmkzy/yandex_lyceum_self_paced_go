package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetMark(name string) (int, error) {
	url := "http://localhost:8082/mark?name=" + url.QueryEscape(name)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("error status")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	mark := strings.TrimSpace(string(body))
	intMark, err := strconv.Atoi(mark)
	if err != nil {
		return 0, err
	}
	return intMark, nil
}
func CompareList(names []string) (map[string]string, error) {
	var res int
	for _, name := range names {
		mark, err := GetMark(name)
		if err != nil {
			return nil, err
		}
		res += mark
	}
	if len(names) == 0 {
		return nil, errors.New("size 0")
	}
	res /= len(names)
	mapa := make(map[string]string)
	for _, name := range names {
		mark, err := GetMark(name)
		if err != nil {
			return nil, err
		}
		if mark == res {
			mapa[name] = "="
		} else if mark > res {
			mapa[name] = ">"
		} else {
			mapa[name] = "<"
		}
	}
	return mapa, nil

}
