package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// StartServer запускает HTTP-сервер
func StartServer(address string) {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Не удалось создать директорию для загрузок: %v", err)
	}

	router := NewRouter()
	fmt.Printf("Сервер запущен на %s\n", address)
	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
