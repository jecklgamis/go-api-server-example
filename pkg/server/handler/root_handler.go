package server

import (
	"encoding/json"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"name": "go-api-server-template", "message": "It works on my machine!"})
}
