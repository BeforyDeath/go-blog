package core

import (
    log "github.com/Sirupsen/logrus"
    "os"
)

func (l *Logger) Init() {
    if Config.Logger.Debug {
        log.SetLevel(log.DebugLevel)
    }
}

func (l *Logger) File() (file *os.File) {
    logFile, err := os.OpenFile("core.log", os.O_CREATE | os.O_RDWR, 0666) //os.O_APPEND |
    if err != nil {
        log.Fatalf("error opening file: %v", err)
    }
    log.SetOutput(logFile)
    log.SetFormatter(&log.JSONFormatter{})
    return logFile
}
