package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"sort"
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
func BestStudents(names []string) (string, error) {
	var res int
	mapa := make(map[string]int)
	for _, name := range names {
		mark, err := GetMark(name)
		if err != nil {
			return "", err
		}
		mapa[name] = mark
		res += mark
	}
	res /= len(names)
	var ans []string
	for _, name := range names {
		mark := mapa[name]
		if mark > res {
			ans = append(ans, name)
		}
	}
	sort.Strings(ans)
	return strings.Join(ans, ","), nil
}
