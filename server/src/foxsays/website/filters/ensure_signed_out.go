package filters

import (
	"foxsays/config"
	"net/http"
)

func EnsureSignedOut(w http.ResponseWriter, r *http.Request) {
	session := config.Session.ForWebsite(r)
	_, signedIn := session.Values["uid"]

	if signedIn {
		delete(session.Values, "uid")
		session.Save(r, w)
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
	}
}
