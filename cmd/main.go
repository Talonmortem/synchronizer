package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"synchronizer/synchronizer-service/server"

	"github.com/Talonmortem/synchronizer/synchronizer-service/client"
)

func main() {
	mode := flag.String("mode", "server", "Режим запуска: 'server' для запуска сервера, 'upload' для загрузки файла, 'download' для скачивания файла")
	filePath := flag.String("file", "", "Путь к файлу для загрузки или имя файла для скачивания")
	serverURL := flag.String("url", "http://localhost:8080", "URL сервера для загрузки/скачивания файлов")
	flag.Parse()

	switch *mode {
	case "server":
		server.StartServer(":8080")
	case "upload":
		if *filePath == "" {
			log.Fatal("Укажите путь к файлу для загрузки с помощью параметра -file")
		}
		if err := client.UploadFile(*serverURL, *filePath); err != nil {
			log.Fatalf("Ошибка загрузки файла: %v", err)
		}
	case "download":
		if *filePath == "" {
			log.Fatal("Укажите имя файла для скачивания с помощью параметра -file")
		}
		if err := client.DownloadFile(*serverURL, *filePath); err != nil {
			log.Fatalf("Ошибка скачивания файла: %v", err)
		}
	default:
		fmt.Println("Неверный режим. Используйте 'server', 'upload' или 'download'")
		os.Exit(1)
	}
}
