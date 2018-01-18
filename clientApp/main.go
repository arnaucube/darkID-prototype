package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
)

const keysDir = "keys"
const keysize = 2048
const hashize = 1536

func main() {
	color.Blue("Starting darkID clientApp")

	readConfig("config.json")
	fmt.Println(config)

	//create models directory
	_ = os.Mkdir(keysDir, os.ModePerm)

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
	//here, run webserver
	log.Println("webserver in port " + "8080")
	http.Handle("/", http.FileServer(http.Dir("./GUI")))
	http.ListenAndServe(":"+"8080", nil)
}
