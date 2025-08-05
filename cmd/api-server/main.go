package main

import (
	"log"
	"net/http"

	"github.com/Akashkarmokar/golang_server_setup/internal/router"
)

func main() {
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
