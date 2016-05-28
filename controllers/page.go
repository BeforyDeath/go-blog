package controllers

import (
	"github.com/BeforyDeath/go-blog/core"
	"github.com/BeforyDeath/go-blog/models"
	log "github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/BeforyDeath/pagination"
	"strconv"
)

type PageController struct {
	Pagination *pagination.Pagination
}

func (pc *PageController) InitPagination(p int) error {
	if pc.Pagination == nil {

		log.Info("Initial Pagination PageController")

		pc.Pagination = pagination.Create(0, core.Config.PageLimit, core.Config.VisibleRange)

		model := models.Pages{}
		count, err := model.GetTotal()
		if err != nil {
			log.Error(err.Error())
			return err
		}

		pc.Pagination.SetTotal(count)
	}

	pc.Pagination.Get(p)
	return nil
}

func (pc *PageController) GetList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	p := 1
	var err error
	if r.FormValue("p") != "" {
		p, err = strconv.Atoi(r.FormValue("p"))
		if err != nil {
			log.Error(err.Error())
			http.Error(w, http.StatusText(400), 400)
			return
		}
	}
	pc.InitPagination(p)

	model := models.Pages{}

	res, err := model.GetList(pc.Pagination.GetOffset(), pc.Pagination.Limit)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	core.Themes.Result["data"] = res
	core.Themes.Result["pages"] = pc.Pagination.Pages
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
