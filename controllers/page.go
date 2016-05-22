package controllers

import (
    log "github.com/Sirupsen/logrus"
    "github.com/beforydeath/go-blog/core"
    "github.com/beforydeath/go-blog/models"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "strconv"
    "time"
)

type PageController struct {
}

func (self *PageController) View(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    model := models.Pages{}
    res, err := model.GetByAlias(ps.ByName("alias"))
    if err != nil {
        log.Info(ps.ByName("alias") + ": " + err.Error())
        http.Error(w, http.StatusText(404), 404)
        return
    }
    // todo reparse template
    core.Themes.Init()

    err = core.Theme.ExecuteTemplate(w, "element", res)
    if err != nil {
        log.Error(err.Error())
        return
    }
}

func (self *PageController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    model := models.Pages{}
    res, err := model.GetList()
    if err != nil {
        log.Error(err.Error())
        http.Error(w, http.StatusText(500), 500)
    }
    // todo reparse template
    core.Themes.Init()

    err = core.Theme.ExecuteTemplate(w, "elementList", res)
    if err != nil {
        log.Error(err.Error())
        return
    }
}

func (self *PageController) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

    // todo reparse template
    core.Themes.Init()

    err = core.Theme.ExecuteTemplate(w, "elementForm", res)
    if err != nil {
        log.Error(err.Error())
        return
    }
}

func (self *PageController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    page := new(models.Page)

    if ( r.Method == "POST") {
        r.ParseForm()
        err := decoder.Decode(page, r.PostForm)
        if err != nil {
            log.Error(err.Error())
        }
        page.Created_at = time.Now()

        model := models.Pages{}
        id, err := model.Create(page)
        if err != nil {
            log.Error(err.Error())
        }
        log.Infof("Create new id:%d", id)
    }

    // todo reparse template
    core.Themes.Init()

    err := core.Theme.ExecuteTemplate(w, "elementForm", page)
    if err != nil {
        log.Error(err.Error())
        return
    }
}

func (self *PageController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
