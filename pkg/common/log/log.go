package log

import (
	"log"
	"os"
	"runtime"

	"github.com/alisyahbana/efishery-test/pkg/common/config"
	"github.com/alisyahbana/efishery-test/pkg/common/env"
)

type Config struct {
	LogPath  string `json:"logPath"`
	LogLevel string `json:"logLevel"`
}

var logLevel int

func init() {
	var cfg *Config

	config.LoadConfiguration(&cfg, "log", env.GetEnv())
	f, err := os.OpenFile(cfg.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("error opening file log")
	}

	log.SetOutput(f)
	logLevel = getLogLevel(cfg.LogLevel)
}

func Debug(message string) {
	if logLevel <= 1 {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("Debug : %s:%d %s\n", fn, line, message)
	}
}

func Info(message string) {
	if logLevel <= 2 {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("Info : %s:%d %s\n", fn, line, message)
	}
}

func Error(message string) {
	if logLevel <= 3 {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("Error : %s:%d %s\n", fn, line, message)
	}
}

func getLogLevel(logLevel string) int {
	logLevelMap := make(map[string]int)
	logLevelMap["DEBUG"] = 1
	logLevelMap["INFO"] = 2
	logLevelMap["ERROR"] = 3
	return logLevelMap[logLevel]
}
