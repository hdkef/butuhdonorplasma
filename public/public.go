package public

import (
	"butuhdonorplasma/dbdriver"
	"butuhdonorplasma/public/add"
	"butuhdonorplasma/public/delete"
	"butuhdonorplasma/public/find"
	"butuhdonorplasma/public/index"
	"butuhdonorplasma/public/result"
	"net/http"
)

type PublicPages struct {
	DBRepo *dbdriver.DBRepo
}

func GetPublicPages(dbrepo *dbdriver.DBRepo) *PublicPages {
	return &PublicPages{
		DBRepo: dbrepo,
	}
}

func (x *PublicPages) AddPage() http.HandlerFunc {
	return add.GetAddHandler(x.DBRepo).Add()
}

func (x *PublicPages) FindPage() http.HandlerFunc {
	return find.GetFindHandler().Find()
}

func (x *PublicPages) IndexPage() http.HandlerFunc {
	return index.GetIndexHandler().Index()
}

func (x *PublicPages) ResultPage() http.HandlerFunc {
	return result.GetResultHandler(x.DBRepo).Result()
}

func (x *PublicPages) DeletePage() http.HandlerFunc {
	return delete.GetDeleteHandler(x.DBRepo).Delete()
}
