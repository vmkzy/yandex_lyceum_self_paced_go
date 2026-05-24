package main

import (
	"errors"
	"os"
)

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("file not open")
	}
	defer file.Close()
	_, err = file.WriteString(message)
	if err != nil {
		return errors.New("message not write")

	}
	return nil
}
