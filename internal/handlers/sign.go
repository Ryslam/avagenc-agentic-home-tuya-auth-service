package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Ryslam/avagenc-agentic-home-tuya-auth-service/internal/models"
	"github.com/Ryslam/avagenc-agentic-home-tuya-auth-service/internal/services"
)

func SignatureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userReq models.SignRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	accessID := os.Getenv("TUYA_ACCESS_ID")
	accessSecret := os.Getenv("TUYA_ACCESS_SECRET")
	baseURL := os.Getenv("TUYA_BASE_URL")

	if accessID == "" || accessSecret == "" || baseURL == "" {
		http.Error(w, "Server configuration error: environment variables not set", http.StatusInternalServerError)
		return
	}

	tokenSignReq := models.SignRequest{
		Method:  "GET",
		URLPath: "/v1.0/token?grant_type=1",
		Body:    "",
	}
	tokenSignature, err := services.GetSign(accessID, accessSecret, tokenSignReq, "")
	if err != nil {
		log.Printf("Error signing token request: %v", err)
		http.Error(w, "Failed to sign token request: "+err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken, err := services.GetAccessToken(accessID, baseURL, tokenSignature)
	if err != nil {
		log.Printf("Error getting access token: %v", err)
		http.Error(w, "Failed to get access token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	finalSignature, err := services.GetSign(accessID, accessSecret, userReq, accessToken)
	if err != nil {
		log.Printf("Error signing final request: %v", err)
		http.Error(w, "Failed to sign final request: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(finalSignature)
}
