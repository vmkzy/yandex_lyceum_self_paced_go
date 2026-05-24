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

func DecodeStudentFromReader(r io.Reader) (Student, error) {
	decoder := json.NewDecoder(r)
	var student Student
	err := decoder.Decode(&student)
	if err != nil {
		return Student{}, errors.New("error")
	}
	return student, nil
}
