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
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=windows-1251")
		w.Header().Set("Content-Disposition", "attachment; filename=example.json")
		http.ServeFile(w, r, "../internal/file/static/cons-layer-api-v1_0-example.json");
	}
}