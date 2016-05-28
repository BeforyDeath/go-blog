package controllers

import (
	"bytes"
	"encoding/base64"
	"github.com/BeforyDeath/go-blog/core"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type UserController struct {
	Name     []byte
	Password []byte
}

func (c *UserController) BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const basicAuthPrefix string = "Basic "
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 &&
					bytes.Equal(pair[0], c.Name) &&
					bytes.Equal(pair[1], c.Password) {
					h(w, r, ps)

					core.Themes.Result["login"] = true

					return
				}
			}
		}

		core.Themes.Result["login"] = false

		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
