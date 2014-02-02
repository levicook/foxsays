package sign_in

import (
	"foxsays/pages"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := pages.GetPage(`pages/sign_in`)
	p.WriteTitle(`Sign In | Foxsays`)
	p.Render(w)
}
