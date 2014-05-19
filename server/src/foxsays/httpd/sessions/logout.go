package sessions

import (
	"foxsays/httpd/status"
	"net/http"
)

func Logout(urlStr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Get(r).Clear(w)
		http.Redirect(w, r, urlStr, status.Found)
	}
}
