package main

import (
	"log"
	"os"
)

var Glf *os.File

func init() {
	logDir := "/home/logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(logDir+"/monitor.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err == nil {
		log.SetOutput(f)
	} else {
		log.Fatal(err)
	}

	Glf = f
}
