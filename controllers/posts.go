package controllers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/beforydeath/go-blog/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Posts struct {
}

func (p *Posts) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	Posts := models.Posts{}
	res, err := Posts.GetList()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	log.Info(res)
}
