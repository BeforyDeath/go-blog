package controllers

import (
    log "github.com/Sirupsen/logrus"
    "github.com/beforydeath/go-blog/core"
    "github.com/beforydeath/go-blog/models"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

type ElementController struct {
}

func (e *ElementController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    Elements := models.Elements{}
    res, err := Elements.GetList()
    if err != nil {
        http.Error(w, http.StatusText(500), 500)
    }

    err = core.Theme.ExecuteTemplate(w, "index", res)
    if err != nil {
        log.Error(err.Error())
        return
    }

}
