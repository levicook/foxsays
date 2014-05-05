package images

import (
	"fmt"
	"net/http"
	"foxsays/mime"
)

func validateWebImage(contentType string) (statusCode int, statusText string) {
	if !mime.IsImage(contentType) {
		statusCode = http.StatusUnsupportedMediaType
		statusText = fmt.Sprintf("Unsupported Media Type: %q", contentType)
	}

	return
}
