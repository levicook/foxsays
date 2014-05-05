package config

import (
	"net/http"
	"time"
)

func expires(h http.Handler) http.Handler {
	oneYear := time.Hour * 24 * 365

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rp := r.URL.Path

		if hasSuffix(rp, ".css") || hasSuffix(rp, ".js") {
			farFuture := time.Now().Add(oneYear).Format(http.TimeFormat)
			w.Header().Add("Expires", farFuture)
		}

		h.ServeHTTP(w, r)
	})
}
