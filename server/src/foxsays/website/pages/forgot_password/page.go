package forgot_password

import (
	"foxsays/pages"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := pages.GetPage(`pages/forgot_password`)
	p.WriteTitle(`Forgot Password | Foxsays`)
	p.Render(w)
}
