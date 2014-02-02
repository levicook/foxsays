package filters

import (
	"foxsays/config"
	"net/http"
	"net/url"
)

func EnsureSignedIn(w http.ResponseWriter, r *http.Request) {
	session := config.Session.ForWebsite(r)
	_, signedIn := session.Values["uid"]

	if !signedIn {
		q := new(url.Values)
		q.Set("revisit", r.RequestURI)

		u := new(url.URL)
		u.Path = "/sign_in"
		u.RawQuery = q.Encode()

		http.Redirect(w, r, u.String(), http.StatusFound)
	}
}
