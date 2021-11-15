package main

import (
	"log"
	"os"

	"github.com/betNevS/code-examples/log/cuslog"
)

func main() {
	cuslog.Info("std log")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))
	cuslog.Debug("change std log to debug level")

	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create failed")
	}
	defer fd.Close()
	l := cuslog.New(cuslog.WithLevel(cuslog.InfoLevel), cuslog.WithOutput(fd), cuslog.WithFormatter(&cuslog.JsonFormatter{}))
	l.Info("cuslog info")

}
