package requests

import (
	"fmt"
	"io"
	"net/http"
)

func DeleteRequest(client *http.Client, url string) error {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("Ошибка при создании запроса: %v \n", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Ошибка при отправке запроса: %v \n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Статус: %d, Ответ: %s\n", resp.StatusCode, string(body))

	return nil
}
