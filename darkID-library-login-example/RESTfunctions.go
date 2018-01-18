package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cryptoballot/rsablind"
	"github.com/fatih/color"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    string        `json:"token"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "clientApp")
}

func Signup(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Print("user signup: ")
	fmt.Println(user)

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

func Login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var key Key
	err := decoder.Decode(&key)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//TODO check if the user password exists in the database

	fmt.Print("key login: ")
	fmt.Println(key)
	token, err := newToken()
	check(err)

	//validate if the pubK darkID is in the blockchain

	//verify that the darkID is signed
	if err := rsablind.VerifyBlindSignature(key.ServerVerifier, key.Hashed, key.UnblindedSig); err != nil {
		fmt.Println(err)
	} else {
		color.Green("blind signature verified")
	}

	/*jResp, err := json.Marshal(token)
	if err != nil {
		panic(err)
	}*/
	fmt.Fprintln(w, string(token))
}
