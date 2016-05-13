package models

import (
    "database/sql"
    log "github.com/Sirupsen/logrus"

    "github.com/beforydeath/go-blog/core"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
    var err error
    db, err = sql.Open(core.Config.DataBase.DriverName, core.Config.DataBase.DataSourceName)
    if err != nil {
        log.Fatal(err.Error())
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }
}
func CloseDB() {
    db.Close()
}
