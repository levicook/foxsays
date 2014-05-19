package dashboard

import (
	"foxsays/httpd/page"
	"foxsays/httpd/sessions"
	"foxsays/httpd/status"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	s := sessions.Get(r)

	if s.RealUserId().Blank() || !s.RealUser().SuperUser {
		http.Redirect(w, r, "/admin", status.Forbidden)
		return
	}

	show(w, r)
}

func show(w http.ResponseWriter, r *http.Request) {
	p := page.Page{Id: `admin/pages/dashboard`}
	p.Render(w)
}
