package delete

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"errors"
	"net/http"
	"path/filepath"
)

type DeleteHandler struct {
	DBRepo *dbdriver.DBRepo
}

var dir string = filepath.Join("public", "delete", "delete.html")
var dir_success string = filepath.Join("public", "delete", "delete-success.html")
var dir_failed string = filepath.Join("public", "utils", "fail.html")

func GetDeleteHandler(dbrepo *dbdriver.DBRepo) *DeleteHandler {
	return &DeleteHandler{
		DBRepo: dbrepo,
	}
}

func (x *DeleteHandler) Delete() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			handleDeletePost(rw, r, x)
			return
		}

		controller.RenderPage(rw, r, nil, dir)
	}
}

func handleDeletePost(rw http.ResponseWriter, r *http.Request, x *DeleteHandler) {
	err := r.ParseForm()
	if err != nil {
		controller.RenderFailed(rw, r, err, dir_failed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		controller.RenderFailed(rw, r, errors.New("NO ID inserted"), dir_failed)
		return
	}

	res, err := x.DBRepo.DeletePatientByID(id)
	if err != nil {
		controller.RenderFailed(rw, r, err, dir_failed)
		return
	}

	controller.RenderPage(rw, r, res, dir_success)
}
