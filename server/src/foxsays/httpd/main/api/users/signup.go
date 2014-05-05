package users

import (
	"encoding/json"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	json.NewEncoder(w).Encode(map[string]string{
		"foo": "bar",
	})
}
