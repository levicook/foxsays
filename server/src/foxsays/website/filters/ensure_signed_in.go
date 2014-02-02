package filters

import (
	"net/http"
)

func EnsureSignedIn(w http.ResponseWriter, r *http.Request) {
	// get the session
	// do they have a user id?
	// if not, redirect them to the "sign_in" route, passing r.RequestURI as a return path
}
