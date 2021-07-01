package find

import (
	"butuhdonorplasma/controller"
	"butuhdonorplasma/dbdriver"
	"embed"
	"html/template"
	"net/http"
)

//go:embed *.html
var tmpl embed.FS
var thistemplates *template.Template = template.Must(template.ParseFS(tmpl, "find.html"))

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

		controller.RenderPage(rw, r, data, thistemplates)
	}
}
