package core

import (
    "html/template"
    "time"
)

var Themes themes

type themes struct {
}

var Theme *template.Template

func (t *themes) Init() {
    funcMap := template.FuncMap{
        "formatTime": func(arg time.Time) string {
            // yyyy-MM-dd HH:mm:ss - "2006-01-02 15:04:05".
            return arg.Format("02-01-2006 15:04");
        },
    }
    Theme = template.Must(template.New("main").Funcs(funcMap).ParseGlob(Config.ThemePath + "/*.hbs"))
}
