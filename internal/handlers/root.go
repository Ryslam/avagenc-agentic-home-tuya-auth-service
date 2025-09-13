package handlers

import (
	"encoding/json"
	"net/http"
)

// RootResponse is the structure for the root endpoint's JSON response
type RootResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

// RootHandler handles requests to the root endpoint ("/")
func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := RootResponse{
		Service: "avagenc-agentic-home-tuya-auth-service",
		Version: "0.1.0", // For now, we'll hardcode this.
		Status:  "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
