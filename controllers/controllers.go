package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/russross/blackfriday"
	"net/http"
)

type Controller struct {
	Page PageController
	User UserController
}

func (c *Controller) Markdown(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	text := r.PostFormValue("text")
	md := blackfriday.MarkdownBasic([]byte(text))
	w.WriteHeader(200)
	fmt.Fprint(w, string(md))
}
