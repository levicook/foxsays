package settings

import (
	"foxsays/config"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := config.Website.GetPage(`pages/settings`)
	p.WriteTitle(`Settings | Foxsays`)
	p.Render(w)
}
