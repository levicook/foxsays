package settings

import (
	"foxsays/pages"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := pages.GetPage(`pages/settings`)
	p.WriteTitle(`Settings | Foxsays`)
	p.Render(w)
}
