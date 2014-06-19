package images

import (
	"foxsays/httpd/status"
	"net/http"
)

func validateContentLength(r *http.Request) (statusCode int, statusText string) {
	s := toSize(r.Header.Get("Content-Length"))

	switch {
	case s <= 0:
		return status.LengthRequired, "Length Required"
	case s > MaxSize: // todo read this from a config file?
		return status.RequestEntityTooLarge, "Request Entity Too Large"
	}

	return
}
