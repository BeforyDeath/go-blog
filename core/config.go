package core

import (
    "encoding/json"
    log "github.com/Sirupsen/logrus"
    "os"
)

var Config config

type config struct {
    BasePath string
    Theme    string
    Logger   Logger
    DataBase Database
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
}
