package images

import (
	"foxsays/config"
	"foxsays/httpd/sessions"
	"foxsays/httpd/status"
	"foxsays/httpd/utils"
	"foxsays/log"
	"foxsays/models"
	"io"
	"net/http"

	"github.com/levicook/head"
)

func Create(w http.ResponseWriter, r *http.Request) {

	if sc, st := utils.ValidateSuperUser(r); sc > 0 {
		http.Error(w, st, sc)
		return
	}

	if sc, st := validateContentLength(r); sc > 0 {
		http.Error(w, st, sc)
		return
	}

	if sc, st := validateMultipartFormData(r); sc > 0 {
		http.Error(w, st, sc)
		return
	}

	fileRepo := config.Repos.NewFileRepo()

	file, err := fileRepo.Create()
	log.PanicIf(err)

	mr, err := r.MultipartReader()
	log.PanicIf(err)

	hw := head.NewWriter(512)

	{ // stream our upload; capture FileName and enforce MaxSize (b/c Content-Length can be a lie)
		lr := new(io.LimitedReader)
		lr.N = int64(MaxSize)
		for {
			if part, err := mr.NextPart(); err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			} else {

				if file.Name() == "" && part.FileName() != "" {
					file.SetName(part.FileName())
				}

				lr.R = part
				io.Copy(file, io.TeeReader(lr, hw))
				part.Close()
			}
		}
	}

	{ // detect and set file content type
		ct := http.DetectContentType(hw.Head())
		file.SetContentType(ct)
		file.Close()
	}

	if sc, st := validateWebImage(file.ContentType()); sc > 0 {
		fileRepo.RemoveId(file.Id())
		http.Error(w, st, sc)
		return
	}

	createdBy := sessions.Get(r).RealUserId()
	image := models.NewImage(file, createdBy)

	{ // save image
		imageRepo := config.Repos.NewImageRepo()
		log.PanicIf(imageRepo.Create(image))
	}

	utils.WriteJson(w, status.OK, image)
}
