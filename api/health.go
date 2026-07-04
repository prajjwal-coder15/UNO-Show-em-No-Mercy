package api

import (
	"encoding/json"
	"net/http"
)

// Health is the health status endpoint of the API.
func Health(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}