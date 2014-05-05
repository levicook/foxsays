package images

import (
	"encoding/json"
	"net/http"
	"foxsays/config"
	"foxsays/httpd/route"
	"foxsays/log"

	"github.com/gorilla/mux"
)

func Meta(w http.ResponseWriter, r *http.Request) {
	meta(w, r, route.Vars(mux.Vars(r)))
}

func meta(w http.ResponseWriter, r *http.Request, v route.Vars) {
	imageRepo := config.Repos.OpenImageRepo()

	image, err := imageRepo.OneById(v.FileId("imageId"))
	log.PanicIf(err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(image)
}
