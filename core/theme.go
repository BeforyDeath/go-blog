package core

import (
    log "github.com/Sirupsen/logrus"
    "html/template"
)

var Themes themes

type themes struct {
}

var Theme *template.Template

func (t *themes) Init() {
    var err error
    Theme, err = template.ParseGlob(Config.ThemePath + "/*.hbs")
    if err != nil {
        log.Error(err.Error())
        return
    }
}
