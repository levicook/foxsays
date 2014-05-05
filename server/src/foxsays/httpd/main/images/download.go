package images

import (
	"fmt"
	"io"
	"net/http"
	"foxsays/config"
	"foxsays/httpd/route"
	"foxsays/log"

	"github.com/gorilla/mux"
)

func Download(w http.ResponseWriter, r *http.Request) {
	download(w, r, route.Vars(mux.Vars(r)))
}

func download(w http.ResponseWriter, r *http.Request, v route.Vars) {
	fileRepo := config.Repos.OpenFileRepo()

	file, err := fileRepo.OneById(v.FileId("imageId"))
	log.PanicIf(err)
	defer file.Close()

	contentDisposition := fmt.Sprintf("attachment; filename=%q", file.Name())

	w.Header().Set("Content-Disposition", contentDisposition)
	w.Header().Set("Content-Type", file.ContentType())
	io.Copy(w, file)
}
