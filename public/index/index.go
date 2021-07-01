package index

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"embed"
	"html/template"
	"net/http"
)

//go:embed *.html
var tmpl embed.FS
var thistemplates *template.Template = template.Must(template.ParseFS(tmpl, "index.html"))

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
		controller.RenderPage(rw, r, nil, thistemplates)
	}
}
