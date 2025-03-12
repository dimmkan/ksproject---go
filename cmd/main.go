package main

import (
	"fmt"
	"ksproject_go/config"
	"ksproject_go/internal/services"
	"ksproject_go/pkg/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found!")
	}
}



func main() {
	router := http.NewServeMux()
	_ = db.NewDb(config.New())
	services.NewServicesHandler(router)
	
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("Server started")
	server.ListenAndServe()
}