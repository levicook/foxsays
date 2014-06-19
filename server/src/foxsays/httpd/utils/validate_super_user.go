package utils

import (
	"foxsays/httpd/sessions"
	"foxsays/httpd/status"
	"net/http"
)

func ValidateSuperUser(r *http.Request) (statusCode int, statusText string) {
	s := sessions.Get(r)

	if !s.RealUser().SuperUser {
		statusCode = status.Unauthorized
		statusText = "Unauthorized"
	}

	return
}
