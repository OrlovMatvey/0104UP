package main

import (
	"net/http"
	"project/requests"
)

func main() {
	client := &http.Client{}
	requests.UploadRequest(client, "http://localhost:8070/upload", "client_files/Image.jpg")
	requests.ListRequest(client, "http://localhost:8070/list")
	requests.DownloadRequest(client, "http://localhost:8070/download/123.txt")
	requests.DeleteRequest(client, "http://localhost:8070/delete/321.txt")
}
