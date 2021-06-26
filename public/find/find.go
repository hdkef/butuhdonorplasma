package find

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"net/http"
	"path/filepath"
)

var dir string = filepath.Join("public", "find", "find.html")

type FindHandler struct {
	DBRepo *dbdriver.DBRepo
}

func GetFindHandler(dbrepo *dbdriver.DBRepo) *FindHandler {
	return &FindHandler{
		DBRepo: dbrepo,
	}
}

func (x *FindHandler) Find() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		data := controller.GetProvince()

		controller.RenderPage(rw, r, data, dir)
	}
}
