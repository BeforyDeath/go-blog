package core

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
)

var Config config

type config struct {
	BasePath       string
	Theme          string
	ThemeReload    bool
	ThemePath      string
	AdminThemePath string
	Logger         Logger
	DataBase       Database
	Listen         string
}

type Logger struct {
	Debug   bool
	OutFile bool
}

type Database struct {
	DriverName     string
	DataSourceName string
}

func (c *config) Init() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	if Config.Listen == "" {
		Config.Listen = ":8085"
	}

	cwd, _ := os.Getwd()
	if Config.BasePath != "" {
		cwd = Config.BasePath
	} else {
		Config.BasePath = cwd
	}
	c.ThemePath = filepath.Join(cwd, "/themes/" + Config.Theme)
	c.AdminThemePath = filepath.Join(cwd, "/themes/admin")
}
