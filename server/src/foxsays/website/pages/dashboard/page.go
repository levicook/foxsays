package dashboard

import (
	"foxsays/pages"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := pages.GetPage(`pages/dashboard`)
	p.WriteTitle(`Dashboard | Foxsays`)
	p.Render(w)
}
