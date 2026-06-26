package routers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Неправильный HTTP метод!", http.StatusMethodNotAllowed)
		return
	}

	filename := filepath.Base(req.URL.Path)
	fullPath := filepath.Join(dataDir, filename)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(res, "Файл не найден!", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Write(data)
}
