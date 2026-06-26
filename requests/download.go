package requests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadRequest(client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Ошибка при отправке запроса: %v \n", err)
	}
	defer resp.Body.Close()

	filename := filepath.Base(url)
	filepath := filepath.Join("client_files", filename)

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Ошибка при создании файла!", err)
	}
	defer file.Close()

	body, _ := io.ReadAll(resp.Body)
	_, err = file.Write(body)
	if err != nil {
		return fmt.Errorf("Ошибка при записи содержимого файла!", err)
	}
	fmt.Printf("Файл успешно скачан!\n")

	return nil
}
