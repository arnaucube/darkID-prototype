package main

import (
	"log"
	"runtime"
)

func check(err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Println(line)
		log.Println(fn)
		log.Println(err)
	}
}
