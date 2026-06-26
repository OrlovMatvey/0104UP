package routers

import (
	"net/http"
	"os"
	"path/filepath"
)


func DeleteFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(res, "Неправильный HTTP метод!", http.StatusMethodNotAllowed)
		return
	}

	filename := filepath.Base(req.URL.Path)
	filepath := filepath.Join(dataDir, filename)

	if err := os.Remove(filepath); err != nil {
		http.Error(res, "Не удалось найти файл!", http.StatusNotFound)
		return
	}

	res.Write([]byte("Файл был успешно удален!"))
}
