package images

import (
	"foxsays/config"
	"foxsays/httpd/route"
	"foxsays/httpd/status"
	"foxsays/httpd/utils"
	"foxsays/log"
	"net/http"

	"github.com/gorilla/mux"
)

func Meta(w http.ResponseWriter, r *http.Request) {
	meta(w, r, route.Vars(mux.Vars(r)))
}

func meta(w http.ResponseWriter, r *http.Request, v route.Vars) {
	imageRepo := config.Repos.NewImageRepo()

	image, err := imageRepo.OneById(v.FileId("imageId"))
	log.PanicIf(err)

	utils.WriteJson(w, status.OK, image)
}
