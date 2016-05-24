package controllers

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/beforydeath/go-blog/core"
	"github.com/beforydeath/go-blog/models"
	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

var decoder = schema.NewDecoder()

func (pc *PageController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page := new(models.Page)

	if r.Method == "POST" {
		r.ParseForm()
		err := decoder.Decode(page, r.PostForm)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err.Error())
			return
		}
		page.Created_at = time.Now()

		model := models.Pages{}
		id, err := model.Create(page)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err.Error())
			return
		}
		w.WriteHeader(201)
		fmt.Fprint(w, id)
		return
	}

	core.Themes.Reload()
	err := core.Theme.ExecuteTemplate(w, "pageForm", page)
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func (pc *PageController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(404), 404)
		return
	}

	model := models.Pages{}
	res, err := model.GetById(id)
	if err != nil {
		log.Info(ps.ByName("id") + ": " + err.Error())
		http.Error(w, http.StatusText(404), 404)
		return
	}

	core.Themes.Reload()
	err = core.Theme.ExecuteTemplate(w, "pageForm", res)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
