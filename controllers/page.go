package controllers

import (
	"github.com/BeforyDeath/go-blog/core"
	"github.com/BeforyDeath/go-blog/models"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PageController struct {
}

func (pc *PageController) GetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	model := models.Pages{}

	err := model.GetTotal()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	model.Pagination.Get(1)

	res, err := model.GetList()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	core.Themes.Result["data"] = res
	core.Themes.Result["pages"] = model.Pagination.Pages
	core.Themes.Result["meta"] = map[string]string{"title": "All page"}

	core.Themes.Reload()
	err = core.Theme.ExecuteTemplate(w, "pageList", core.Themes.Result)
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func (pc *PageController) GetByAlias(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	model := models.Pages{}
	res, err := model.GetByAlias(ps.ByName("alias"))
	if err != nil {
		log.Info(ps.ByName("alias") + ": " + err.Error())
		http.Error(w, http.StatusText(404), 404)
		return
	}

	core.Themes.Result["data"] = res
	core.Themes.Result["meta"] = map[string]string{"title": res.Name}

	core.Themes.Reload()
	err = core.Theme.ExecuteTemplate(w, "page", core.Themes.Result)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
