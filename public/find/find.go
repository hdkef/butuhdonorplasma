package find

import (
	"butuhdonorplasma/controller"
	"embed"
	"html/template"
	"net/http"
)

//go:embed *.html
var tmpl embed.FS
var thistemplates *template.Template = template.Must(template.ParseFS(tmpl, "find.html"))

type FindHandler struct{}

func GetFindHandler() *FindHandler {
	return &FindHandler{}
}

func (x *FindHandler) Find() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		data := controller.GetProvince()

		controller.RenderPage(rw, r, data, thistemplates)
	}
}
