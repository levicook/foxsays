package images

import (
	"fmt"
	"foxsays/httpd/status"
	"net/http"
	"strings"
)

func validateMultipartFormData(r *http.Request) (statusCode int, statusText string) {
	t := r.Header.Get("Content-Type")

	if !strings.HasPrefix(t, "multipart/form-data;") {
		statusCode = status.UnsupportedMediaType
		statusText = fmt.Sprintf("Unsupported Media Type: %q", t)
	}

	return
}
