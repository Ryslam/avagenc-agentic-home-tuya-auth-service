package main

import (
	"log"
	"net/http"

	"github.com/Ryslam/avagenc-agentic-home-tuya-auth-service/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Reading credentials from environment.")
	}

	http.HandleFunc("/sign", handlers.SignatureHandler)

	port := ":8080"
	log.Printf("Starting Avagenc Tuya Auth Service on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
