package public_home

import (
	"fmt"
	"foxsays/pages"
	"net/http"
	"time"
)

func Page(w http.ResponseWriter, r *http.Request) {
	p := pages.GetPage(`pages/public_home`)
	p.WriteTitle(`Home | Foxsays`)
	p.WriteMain(fmt.Sprintf("%s", time.Now()))
	p.Render(w)
}
