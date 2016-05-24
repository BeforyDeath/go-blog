package core

import (
	"github.com/russross/blackfriday"
	"html/template"
	"strings"
	"time"
)

var Themes themes
var Theme *template.Template

type themes struct {
	Result map[string]interface{}
}

func (t *themes) Init() {
	funcMap := template.FuncMap{
		"formatTime": func(arg time.Time) string {
			return arg.Format("02-01-2006 15:04") // yyyy-MM-dd HH:mm:ss - "2006-01-02 15:04:05".
		},
		"unescape": func(str string) template.HTML {
			return template.HTML(str)
		},
		"replace": func(arg ...string) string {
			return strings.Replace(arg[1], arg[0], "", -1)
		},
		"markdown": func(str string) template.HTML {
			return template.HTML(blackfriday.MarkdownBasic([]byte(str)))
		},
	}

	Theme = template.Must(template.New("main").Funcs(funcMap).ParseGlob(Config.ThemePath + "/*.html"))

	t.InitResult()
}

func (t *themes) InitResult() {
	if t.Result == nil {
		t.Result = make(map[string]interface{})
		t.Result["login"] = false
	}
}

func (t *themes) Reload() {
	if Config.ThemeReload {
		t.Init()
	}
}
