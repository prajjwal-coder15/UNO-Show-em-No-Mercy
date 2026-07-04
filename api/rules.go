package api

import (
	"encoding/json"
	"net/http"

	"uno/rules"
)

// Rules writes the available rules as JSON.
func Rules(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(rules.Available())
}