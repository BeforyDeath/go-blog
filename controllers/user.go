package controllers

import (
    "bytes"
    "encoding/base64"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "strings"
)

type UserController struct {
    Name     []byte
    Password []byte
}

func (self *UserController) BasicAuth(h httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        const basicAuthPrefix string = "Basic "
        auth := r.Header.Get("Authorization")
        if strings.HasPrefix(auth, basicAuthPrefix) {
            payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
            if err == nil {
                pair := bytes.SplitN(payload, []byte(":"), 2)
                if len(pair) == 2 &&
                bytes.Equal(pair[0], self.Name) &&
                bytes.Equal(pair[1], self.Password) {
                    h(w, r, ps)
                    return
                }
            }
        }
        w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
        http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
    }
}
