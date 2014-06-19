package utils

import (
	"encoding/json"
	"foxsays/log"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	log.PanicIf(json.NewEncoder(w).Encode(body))
}
