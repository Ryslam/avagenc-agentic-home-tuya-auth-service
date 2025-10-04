package main

import (
	"log"
	"net/http"
	"os"

	"github.com/avagenc/agentic-tuya-sign-service/internal/handlers"
	"github.com/joho/godotenv"
)

func authMiddleware(apiKey string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientAPIKey := r.Header.Get("x-avagenc-api-key")
		if clientAPIKey != apiKey {
			log.Printf("Authentication failed: Invalid API Key. Request from %s", r.RemoteAddr)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found. Reading credentials from environment.")
	}

	apiKey := os.Getenv("AVAGENC_API_KEY")
	if apiKey == "" {
		log.Fatal("AVAGENC_API_KEY environment variable not set. Service cannot start.")
	}

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/sign", authMiddleware(apiKey, handlers.SignatureHandler))

	port := ":8080"
	log.Printf("Starting Avagenc Tuya Auth Service on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
