package handlers

import (
	"encoding/json"
	"net/http"
)

func CreateJobAppHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Hello world!",
	})
}
