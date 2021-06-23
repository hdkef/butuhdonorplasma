package public

import (
	"butuhdonorplasma/public/add"
	"butuhdonorplasma/public/find"
	"butuhdonorplasma/public/index"
	"butuhdonorplasma/public/result"
	"net/http"
)

type PublicPages struct {
}

func GetPublicPages() *PublicPages {
	return &PublicPages{}
}

func (x *PublicPages) AddPage() http.HandlerFunc {
	return add.GetAddHandler().Add()
}

func (x *PublicPages) FindPage() http.HandlerFunc {
	return find.GetFindHandler().Find()
}

func (x *PublicPages) IndexPage() http.HandlerFunc {
	return index.GetIndexHandler().Index()
}

func (x *PublicPages) ResultPage() http.HandlerFunc {
	return result.GetResultHandler().Result()
}
