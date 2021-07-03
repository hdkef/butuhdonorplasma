package index

import (
	"butuhdonorplasma/controller"
	"embed"
	"html/template"
	"net/http"
)

//go:embed *.html
var tmpl embed.FS
var thistemplates *template.Template = template.Must(template.ParseFS(tmpl, "index.html"))

type IndexHandler struct {
}

func GetIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (x *IndexHandler) Index() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		controller.RenderPage(rw, r, nil, thistemplates)
	}
}
