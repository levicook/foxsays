package images

import (
	"foxsays/config"
	"foxsays/httpd/route"
	"foxsays/log"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Show(w http.ResponseWriter, r *http.Request) {
	show(w, r, route.Vars(mux.Vars(r)))
}

func show(w http.ResponseWriter, r *http.Request, v route.Vars) {
	fileRepo := config.Repos.NewFileRepo()

	file, err := fileRepo.OneById(v.FileId("imageId"))
	log.PanicIf(err)
	defer file.Close()

	w.Header().Set("Content-Type", file.ContentType())
	io.Copy(w, file)
}
