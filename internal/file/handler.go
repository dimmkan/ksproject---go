package file

import (
	"fmt"
	"ksproject_go/config"
	"net/http"
	"os"
)

type FileHandlerDeps struct{
	Config *config.Config
}

type FileHandler struct{
	Config *config.Config
}

func NewFileHandler(router *http.ServeMux, deps FileHandlerDeps) {
	handler := &FileHandler{
		Config: deps.Config,
	}
	router.HandleFunc("/file", handler.getFile())
}

func (handler *FileHandler) getFile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=windows-1251")
		w.Header().Set("Content-Disposition", "attachment; filename=example.json")
		http.ServeFile(w, r, fmt.Sprintf("%scons-layer-api-v1_0-example.json", handler.Config.FileStorage.STATIC_FILES_STORAGE));
	}
}