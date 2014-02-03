package filters

import (
	"foxsays/config"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	session := config.Session.ForWebsite(r)
	delete(session.Values, "uid")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
