package dashboard

import (
	"net/http"
	"foxsays/httpd/page"
)

func Show(w http.ResponseWriter, r *http.Request) {
	p := page.Page{
		Id:   `main/pages/dashboard`,
		Tail: `<script id="dataPool" type="application/json"></script>`,
	}

	p.Render(w)
}
