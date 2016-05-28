package controllers

import (
	"fmt"
	"github.com/BeforyDeath/go-blog/core"
	"github.com/BeforyDeath/go-blog/models"
	log "github.com/Sirupsen/logrus"
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
		_, err = model.Create(page)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err.Error())
			return
		}

		pc.Pagination.SetTotal(pc.Pagination.Total + 1)
		w.WriteHeader(201)
		fmt.Fprint(w, page.Alias)
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

	if r.Method == "POST" {
		page := new(models.Page)
		r.ParseForm()
		err := decoder.Decode(page, r.PostForm)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err.Error())
			return
		}

		model := models.Pages{}
		_, err = model.Update(page)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err.Error())
			return
		}

		w.WriteHeader(201)
		fmt.Fprint(w, page.Alias)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}

	model := models.Pages{}
	res, err := model.GetById(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}

	core.Themes.Reload()
	err = core.Theme.ExecuteTemplate(w, "pageForm", res)
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func (pc *PageController) Delete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}

	model := models.Pages{}
	res, err := model.Delete(id)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}

	pc.Pagination.SetTotal(pc.Pagination.Total - 1)
	w.WriteHeader(200)
	fmt.Fprint(w, res)
}
