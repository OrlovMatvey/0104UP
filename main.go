package main

import (
	"log"
	"net/http"
	"os"
	"project/routers"
)

const dataDir = "files"

func main() {
	if _, err := os.Stat(dataDir); err != nil {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/upload", routers.UploadFile)
	http.HandleFunc("/download/", routers.DownloadFile)
	http.HandleFunc("/list", routers.ListFiles)
	http.HandleFunc("/delete/", routers.DeleteFile)
	http.ListenAndServe(":8070", nil)
}
