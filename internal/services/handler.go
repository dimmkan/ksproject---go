package services

import (
	"net/http"
)

type ServicesHandler struct{}

func NewServicesHandler(router *http.ServeMux) {
	handler := &ServicesHandler{}
	router.HandleFunc("/services", handler.getAllServices())
}

func (handler *ServicesHandler) getAllServices() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
	}
}
