package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
)

func main() {
	color.Blue("Starting darkID clientApp")

	readConfig("config.json")
	fmt.Println(config)

	//run thw webserver
	go GUI()

	//run API
	log.Println("api server running")
	log.Print("port: ")
	log.Println(config.Port)
	router := NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+config.Port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func GUI() {
	//here, run electron app
}
