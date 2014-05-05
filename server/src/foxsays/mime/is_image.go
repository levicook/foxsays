package mime

import "strings"

var hasPrefix = strings.HasPrefix

func IsImage(contentType string) bool {
	switch {

	case
		hasPrefix(contentType, "image/gif"),
		hasPrefix(contentType, "image/jpeg"),
		hasPrefix(contentType, "image/pjpeg"),
		hasPrefix(contentType, "image/png"):
		return true
	}

	return false
}
