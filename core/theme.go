package core

import (
    log "github.com/Sirupsen/logrus"
    "html/template"
    "os"
    "path/filepath"
)

var Themes themes

type themes struct {
}

var Theme *template.Template

func (t *themes) Init() {
    cwd, _ := os.Getwd()
    if Config.BasePath != "" {
        cwd = Config.BasePath
    }
    var err error
    Theme, err = template.ParseGlob(filepath.Join(cwd, "/themes/" + Config.Theme + "/*.html"))
    if err != nil {
        log.Error(err.Error())
        return
    }
}
