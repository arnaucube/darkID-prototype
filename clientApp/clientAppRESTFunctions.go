package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

//TODO use rsa library instead own rsa functions

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "clientApp")
}

func GetServer(w http.ResponseWriter, r *http.Request) {
	color.Green(config.Server)
	fmt.Fprintln(w, config.Server)
}
func GetIDs(w http.ResponseWriter, r *http.Request) {
	keys := IDs()

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func GetNewID(w http.ResponseWriter, r *http.Request) {
	key := NewID()

	fmt.Println(key)

	jResp, err := json.Marshal(key)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func GetBlindAndSendToSign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idKey := vars["idKey"]
	color.Green(idKey)

	keys := BlindAndSendToSign(idKey)

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

func GetVerify(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packPubK := vars["pubK"]
	color.Green(packPubK)

	//keys := Verify(packPubK)
	keys := "a"

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
