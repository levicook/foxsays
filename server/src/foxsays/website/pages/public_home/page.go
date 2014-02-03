package public_home

import (
	"foxsays/config"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	// if signed in, lookup and redirect to dashboard route

	p := config.Website.GetPage(`pages/public_home`)
	p.WriteTitle(`Home | Foxsays`)
	p.Render(w)
}
