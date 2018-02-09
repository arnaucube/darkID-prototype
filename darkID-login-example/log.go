package main

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func savelog() {
	timeS := time.Now().String()
	_ = os.Mkdir("logs", os.ModePerm)
	//next 3 lines are to avoid windows filesystem errors
	timeS = strings.Replace(timeS, " ", "_", -1)
	timeS = strings.Replace(timeS, ".", "-", -1)
	timeS = strings.Replace(timeS, ":", "-", -1)
	logFile, err := os.OpenFile("logs/log-"+timeS+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
}
