package requests

import (
	"fmt"
	"io"
	"net/http"
)

func ListRequest(client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("Ошибка при отправке запроса: %v \n", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Статус: %d, Ответ: %s\n", resp.StatusCode, string(body))

	return nil
}
