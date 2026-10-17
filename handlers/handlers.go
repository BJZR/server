package handlers

import (
	"encoding/json"
	"net/http"
)

// writeJSON is a helper that writes a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// Home handles GET /
// Returns a welcome/status JSON message.
func Home(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "Welcome to the API",
		"status":  "running",
	})
}

// Health handles GET /health
// Returns {"status":"ok"} with HTTP 200.
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
