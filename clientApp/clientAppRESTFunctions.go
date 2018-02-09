package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

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
func GetID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyID := vars["keyid"]
	key := getKeyByKeyID(keyID)

	jResp, err := json.Marshal(key)
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
	idKey := vars["keyid"]
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

func GetDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyID := vars["keyid"]

	keys := Delete(keyID)

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

type EncryptData struct {
	M string `json:"m"`
	C []byte `json:"c"`
}

func PostEncrypt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyID := vars["keyid"]

	//get ciphertext from POST json
	decoder := json.NewDecoder(r.Body)
	var encryptData EncryptData
	err := decoder.Decode(&encryptData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	encryptData = Encrypt(keyID, encryptData)

	jResp, err := json.Marshal(encryptData)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func PostDecrypt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyID := vars["keyid"]

	//get ciphertext from POST json
	decoder := json.NewDecoder(r.Body)
	var encryptData EncryptData
	err := decoder.Decode(&encryptData)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	encryptData = Decrypt(keyID, encryptData)

	jResp, err := json.Marshal(encryptData)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
