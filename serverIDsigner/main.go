package main

import (
	"fmt"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"

	ownrsa "./ownrsa"
)

var userCollection *mgo.Collection

var serverRSA ownrsa.RSA

func main() {
	color.Blue("Starting serverIDsigner")

	//read configuration file
	readConfig("config.json")

	initializeToken()

	//initialize RSA
	serverRSA = ownrsa.GenerateKeyPair()
	color.Blue("Public Key:")
	fmt.Println(serverRSA.PubK)
	color.Green("Private Key:")
	fmt.Println(serverRSA.PrivK)

	//mongodb
	session, err := getSession()
	check(err)
	userCollection = getCollection(session, "users")

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
