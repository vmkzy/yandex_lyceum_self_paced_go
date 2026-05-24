package main

import (
	"encoding/json"
	"errors"
)

func DeserializeStringMap(data string) (map[string]string, error) {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		return nil, errors.New("error")
	}
	return res, nil
}
