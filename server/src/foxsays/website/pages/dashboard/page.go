package dashboard

import (
	"foxsays/config"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := config.Website.GetPage(`pages/dashboard`)
	p.WriteTitle(`Dashboard | Foxsays`)
	p.Render(w)
}
