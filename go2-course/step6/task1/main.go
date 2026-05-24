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
		return 0, errors.New("Error status")

	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	mark := strings.TrimSpace(string(body))
	intmark, err := strconv.Atoi(mark)
	if err != nil {
		return 0, err
	}
	return intmark, nil
}
func Compare(name1, name2 string) (string, error) {
	mark1, err := GetMark(name1)
	if err != nil {
		return "", err
	}
	mark2, err := GetMark(name2)
	if err != nil {
		return "", err
	}
	if mark1 == mark2 {
		return "=", nil
	} else if mark1 > mark2 {
		return ">", nil
	} else {
		return "<", nil
	}
}
