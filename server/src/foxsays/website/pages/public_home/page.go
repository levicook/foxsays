package public_home

import (
	"foxsays/pages"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	// if signed in, lookup and redirect to dashboard route

	p := pages.GetPage(`pages/public_home`)
	p.WriteTitle(`Home | Foxsays`)
	p.Render(w)
}
