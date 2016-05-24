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
	core.Themes.Init()

	core.Config.Logger.Init()
	if core.Config.Logger.OutFile {
		logFile := core.Config.Logger.File()
		defer logFile.Close()
	}

	models.ConnectDB()
	defer models.CloseDB()

	controller := controllers.Controller{}

	controller.User.Name = []byte("admin")
	controller.User.Password = []byte("password")

	router := httprouter.New()

	router.ServeFiles("/vendor/*filepath", http.FileSystem(http.Dir(core.Config.BasePath+"/themes/assets/vendors/")))
	router.ServeFiles("/assets/*filepath", http.FileSystem(http.Dir(core.Config.ThemePath+"/assets/")))

	router.GET("/", controller.Page.GetList)
	router.GET("/page/:alias", controller.Page.GetByAlias)

	router.GET("/admin/page/create", controller.Page.Create)
	router.POST("/admin/page/create", controller.Page.Create)
	router.GET("/admin/page/update/:id", controller.Page.Update)
	router.POST("/admin/page/update/:id", controller.Page.Update)

	log.Info("Server started ...")
	log.Fatal(http.ListenAndServe(":8085", router))
}
