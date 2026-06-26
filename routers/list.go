package routers

import (
	"fmt"
	"net/http"
	"os"
)

const dataDir = "files"

func ListFiles(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Неправильный HTTP метод!", http.StatusMethodNotAllowed)
		return
	}

	files, err := os.ReadDir(dataDir)
	if err != nil {
		http.Error(res, "Ошибка чтения дериктории!", http.StatusInternalServerError)
		return
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}

	fmt.Fprintf(res, `Файлы на сервере: %q`, names)
}
