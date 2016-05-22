package controllers

import "github.com/gorilla/schema"

type Controller struct {
    Page PageController
    User UserController
}

var decoder = schema.NewDecoder()
