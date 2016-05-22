package controllers

import (
    log "github.com/Sirupsen/logrus"
    "github.com/beforydeath/go-blog/core"
    "github.com/beforydeath/go-blog/models"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "html/template"
)

type PageController struct {
}

func (self *PageController) View(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    element := models.Pages{}
    res, err := element.GetByAlias(ps.ByName("alias"))
    if err != nil {
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
    Elements := models.Pages{}
    res, err := Elements.GetList()
    if err != nil {
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

func (self *PageController) Edit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    tpl, err := template.ParseGlob(core.Config.AdminThemePath + "/*.html")
    if err != nil {
        log.Error(err.Error())
        return
    }
    err = tpl.ExecuteTemplate(w, "elementForm", nil)
}

func (self *PageController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    tpl, err := template.ParseGlob(core.Config.AdminThemePath + "/*.html")
    if err != nil {
        log.Error(err.Error())
        return
    }
    err = tpl.ExecuteTemplate(w, "elementForm", nil)
}

func (self *PageController) Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}