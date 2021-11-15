package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	l := logger.Sugar()
	l.Panicw("msgxxxx", "aa", "bb", "cc", "dd")
	log.Println()
}
