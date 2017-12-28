package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
)

const keysDir = "keys"
const keysize = 2048
const hashize = 1536

var userCollection *mgo.Collection

var serverKey *rsa.PrivateKey

func main() {
	color.Blue("Starting serverIDsigner")

	//read configuration file
	readConfig("config.json")

	//create models directory
	_ = os.Mkdir(keysDir, os.ModePerm)

	initializeToken()

	//initialize RSA
	//generate RSA keys pair
	reader := rand.Reader
	k, err := rsa.GenerateKey(reader, keysize)
	check(err)
	serverKey = k

	savePEMKey(keysDir+"/server_private.pem", k)
	savePublicPEMKey(keysDir+"/server_public.pem", k.PublicKey)

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
