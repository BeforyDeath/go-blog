package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/beforydeath/go-blog/core"
	"github.com/beforydeath/go-blog/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PageController struct {
}

func (pc *PageController) GetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	model := models.Pages{}
	res, err := model.GetList()
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

	core.Themes.Reload()
	err = core.Theme.ExecuteTemplate(w, "pageList", res)
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

	core.Themes.Reload()
	err = core.Theme.ExecuteTemplate(w, "page", res)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
