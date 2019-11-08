package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	//init
	log.SetFormatter(&log.TextFormatter{})
	f, _ := os.Open("../go_study/logrus/log.log")
	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)

	log.Error("test log Error")
	log.Info("test log Info")
}
