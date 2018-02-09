package main

import (
	"io"
	"log"
	"os"
	"time"
)

func savelog() {
	timeS := time.Now().String()
	_ = os.Mkdir("logs", os.ModePerm)
	logFile, err := os.OpenFile("logs/log-"+timeS+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
