package controllers

import (
    log "github.com/Sirupsen/logrus"
    "github.com/beforydeath/go-blog/core"
    "github.com/beforydeath/go-blog/models"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "html/template"
)

type ElementController struct {
}

func (e *ElementController) One(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    //Elements := models.Elements{}

    err := core.Theme.ExecuteTemplate(w, "element", nil)
    if err != nil {
        log.Error(err.Error())
        return
    }
}

func (e *ElementController) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    Elements := models.Elements{}
    res, err := Elements.GetList()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
    }
    err = core.Theme.ExecuteTemplate(w, "elementList", res)
    if err != nil {
        log.Error(err.Error())
        return
    }
}

func (e *ElementController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

    tpl, err := template.ParseGlob(core.Config.AdminThemePath + "/*.hbs")
    if err != nil {
        log.Error(err.Error())
        return
    }
    err = tpl.ExecuteTemplate(w, "elementForm", nil)
}
