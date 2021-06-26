package index

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"net/http"
	"path/filepath"
)

var dir string = filepath.Join("public", "index", "index.html")

type IndexHandler struct {
	DBRepo *dbdriver.DBRepo
}

func GetIndexHandler(dbrepo *dbdriver.DBRepo) *IndexHandler {
	return &IndexHandler{
		DBRepo: dbrepo,
	}
}

func (x *IndexHandler) Index() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		controller.RenderPage(rw, r, nil, dir)
	}
}
