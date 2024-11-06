package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile отправляет файл на сервер
func UploadFile(serverURL, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return fmt.Errorf("не удалось создать часть формы: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("не удалось скопировать файл в форму: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("POST", serverURL+"/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка отправки запроса: %v", err)
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Ответ сервера:", string(resBody))
	return nil
}

// DownloadFile скачивает файл с сервера
func DownloadFile(serverURL, fileName string) error {
	url := fmt.Sprintf("%s/download?file=%s", serverURL, fileName)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("ошибка запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("сервер вернул ошибку: %s", resp.Status)
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %v", err)
	}

	fmt.Printf("Файл %s успешно загружен\n", fileName)
	return nil
}
