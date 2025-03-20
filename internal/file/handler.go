package file

import (
	"net/http"
)

type FileHandler struct{}

func NewFileHandler(router *http.ServeMux) {
	handler := &FileHandler{}
	router.HandleFunc("/file", handler.getFile())
}

func (handler *FileHandler) getFile() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./static/cons-layer-api-v1_0-example.json");
	}
}