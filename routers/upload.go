package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)


func UploadFile(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Неправильный HTTP метод!", http.StatusMethodNotAllowed)
		return
	}

	filename := req.Header.Get("File_Name")
	if filename == "" {
		filename = fmt.Sprintf("File %v", time.Now())
	}

	filename = filepath.Base(filename)
	filepath := filepath.Join(dataDir, filename)

	file, err := os.Create(filepath)
	if err != nil {
		http.Error(res, "Ошибка при создании файла!", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, req.Body)
	if err != nil {
		http.Error(res, "Ошибка при записи содержимого файла!", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("Файл успешно загружен!"))
}
