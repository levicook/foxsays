package images

import (
	"fmt"
	"foxsays/httpd/status"
	"foxsays/mime"
)

func validateWebImage(contentType string) (statusCode int, statusText string) {
	if !mime.IsImage(contentType) {
		statusCode = status.UnsupportedMediaType
		statusText = fmt.Sprintf("Unsupported Media Type: %q", contentType)
	}

	return
}
