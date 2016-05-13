package main

import (
    log "github.com/Sirupsen/logrus"

    "github.com/julienschmidt/httprouter"
    "net/http"

    "github.com/beforydeath/go-blog/controllers"
    "github.com/beforydeath/go-blog/core"
    "github.com/beforydeath/go-blog/models"
)

func main() {
    core.Config.Init()

    core.Config.Logger.Init()
    if core.Config.Logger.OutFile {
        logFile := core.Config.Logger.File()
        defer logFile.Close()
    }

    models.ConnectDB()
    defer models.CloseDB()

    controller := controllers.Controller{}

    router := httprouter.New()
    router.GET("/", controller.Posts.Index)
    log.Fatal(http.ListenAndServe(":8085", router))
}
