package requests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadRequest(client *http.Client, url, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при открытии файла: %v \n", err)
	}
	defer file.Close()

	req, err := http.NewRequest(http.MethodPost, url, file)
	if err != nil {
		return fmt.Errorf("Ошибка при создании запроса: %v \n", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("File_Name", filepath.Base(filePath))

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Ошибка при отправке запроса: %v \n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Статус: %d, Ответ: %s\n", resp.StatusCode, string(body))

	return nil
}
