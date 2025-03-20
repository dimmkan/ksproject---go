package file

import (
	"net/http"
)

type ServicesHandler struct{}

func NewServicesHandler(router *http.ServeMux) {
	handler := &ServicesHandler{}
	router.HandleFunc("/file", handler.getAllServices())
}

func (handler *ServicesHandler) getAllServices() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./static/cons-layer-api-v1_0-example.json");
	}
}