package images

import "net/http"

func validateContentLength(r *http.Request) (statusCode int, statusText string) {
	s := toSize(r.Header.Get("Content-Length"))

	switch {
	case s <= 0:
		return http.StatusLengthRequired, "Length Required"
	case s > MaxSize: // todo read this from a config file?
		return http.StatusRequestEntityTooLarge, "Request Entity Too Large"
	}

	return
}
