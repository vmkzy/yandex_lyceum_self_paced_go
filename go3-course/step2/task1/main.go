package main

import (
	"encoding/json"
	"errors"
)

func SerializeIntSlice(nums []int) ([]byte, error) {
	data, err := json.Marshal(nums)
	if err != nil {
		return nil, errors.New("error")

	}
	return data, err
}
