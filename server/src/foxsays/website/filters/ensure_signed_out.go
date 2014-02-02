package filters

import (
	"net/http"
)

func EnsureSignedOut(w http.ResponseWriter, r *http.Request) {
	// get the session
	// do they have a user id?
	// if so, sign them out and redirect them to r.RequestURI
}
