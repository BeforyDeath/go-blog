package core

import (
    "html/template"
    "github.com/russross/blackfriday"
    "time"
    "strings"
)

var Themes themes

type themes struct {
}

var Theme *template.Template

func (t *themes) Init() {
    funcMap := template.FuncMap{
        "formatTime": func(arg time.Time) string {
            return arg.Format("02-01-2006 15:04") // yyyy-MM-dd HH:mm:ss - "2006-01-02 15:04:05".
        },
        "unescape": func(str string) template.HTML {
            return template.HTML(str)
        },
        "replace": func(arg ...string) string{
            return strings.Replace(arg[1], arg[0], "", -1)
        },
        "markdown": func(str string) template.HTML {
            return template.HTML(blackfriday.MarkdownBasic([]byte(str)))
        },
    }
    Theme = template.Must(template.New("main").Funcs(funcMap).ParseGlob(Config.ThemePath + "/*.html"))
}
