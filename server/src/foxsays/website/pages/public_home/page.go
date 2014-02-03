package public_home

import (
	"foxsays/config"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	session := config.Session.ForWebsite(r)

	if _, signedIn := session.Values["uid"]; signedIn {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		p := config.Website.GetPage(`pages/public_home`)
		p.WriteTitle(`Home | Foxsays`)
		p.Render(w)
	}
}
