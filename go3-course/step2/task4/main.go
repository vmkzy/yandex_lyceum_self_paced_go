package main

import (
	"encoding/json"
	"errors"
	"io"
)

type Student struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func EncodeStudentsToWriter(w io.Writer, students []Student) error {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(students)
	if err != nil {
		return errors.New("error")
	}
	return nil
}
