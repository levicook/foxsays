package sign_in

import (
	"foxsays/config"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := config.Website.GetPage(`pages/sign_in`)
	p.WriteTitle(`Sign In | Foxsays`)
	p.Render(w)
}
