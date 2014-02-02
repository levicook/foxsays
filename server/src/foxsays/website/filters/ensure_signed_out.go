package filters

import (
	"foxsays/config"
	"net/http"
)

func EnsureSignedOut(w http.ResponseWriter, r *http.Request) {
	session := config.Session.ForWebsite(r)

	if !session.IsNew {
		session.Values = make(map[interface{}]interface{})
	}

	session.Save(r, w)
	http.Redirect(w, r, r.RequestURI, http.StatusFound)
}
