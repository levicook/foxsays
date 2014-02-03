package forgot_password

import (
	"foxsays/config"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := config.Website.GetPage(`pages/forgot_password`)
	p.WriteTitle(`Forgot Password | Foxsays`)
	p.Render(w)
}
