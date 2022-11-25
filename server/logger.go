package server

import (
	"log"
	"time"
)

func LogInfo(status int, path string, start time.Time) {
	log.Printf("\033[42m %d \033[0m | PATH: \033[33m\"%s\"\033[0m | DURATION: \033[42m %v \033[0m",
		status, path, time.Since(start))
}

func LogError(status int, path string, err error) {
	log.Printf("\033[41m %d \033[0m | PATH: \033[33m\"%s\"\033[0m | ERROR: \033[31m %v \033[0m",
		status, path, err)
}
