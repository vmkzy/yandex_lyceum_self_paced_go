package main

import (
	"log"
	"os"
)

func main() {
	// Открываем файл
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	// Закрываем файл после выхода из main
	defer file.Close()
	// Конфигурируем логгер, чтобы он выводил лог в файл
	log.SetOutput(file)

	log.Println("Hello, World!")
}
