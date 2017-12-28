package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ownrsa "./ownrsa"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

//TODO use rsa library instead own rsa functions

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "serverIDsigner")
}

func GetServer(w http.ResponseWriter, r *http.Request) {
	color.Green(config.Server)
	fmt.Fprintln(w, config.Server)
}
func IDs(w http.ResponseWriter, r *http.Request) {
	//read the keys stored in /keys directory
	keys := readKeys("keys.json")
	saveKeys(keys, "keys.json")

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func NewID(w http.ResponseWriter, r *http.Request) {
	//generate RSA keys pair
	newKey := ownrsa.GenerateKeyPair()

	key := ownrsa.PackKey(newKey)
	key.Date = time.Now()
	fmt.Println(key)

	keys := readKeys("keys.json")
	keys = append(keys, key)
	saveKeys(keys, "keys.json")

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

type AskBlindSign struct {
	M string `json:"m"`
}

func BlindAndSendToSign(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packPubK := vars["pubK"]
	color.Green(packPubK)

	//read the keys stored in /keys directory
	keys := readKeys("keys.json")

	var key ownrsa.RSA
	//search for complete key
	for _, k := range keys {
		if k.PubK == packPubK {
			key = ownrsa.UnpackKey(k)
		}
	}
	//blind the key.PubK
	var m []int
	//convert packPubK to []bytes
	mBytes := []byte(packPubK)
	for _, byte := range mBytes {
		m = append(m, int(byte))
	}
	rVal := 101
	blinded := ownrsa.Blind(m, rVal, key.PubK, key.PrivK)
	fmt.Println(blinded)

	//convert blinded to string
	var askBlindSign AskBlindSign
	askBlindSign.M = ownrsa.ArrayIntToString(blinded, "_")

	//send to the serverIDsigner the key.PubK blinded
	color.Green(askBlindSign.M)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(askBlindSign)
	res, err := http.Post(config.Server+"blindsign", "application/json", body)
	check(err)
	fmt.Println(res)

	decoder := json.NewDecoder(res.Body)
	//var sigmaString string
	err = decoder.Decode(&askBlindSign)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Println("sigmaString")
	fmt.Println(askBlindSign)
	sigma := ownrsa.StringToArrayInt(askBlindSign.M, "_")
	fmt.Println(sigma)

	//get the serverIDsigner pubK
	serverPubK := getServerPubK(config.Server)

	//unblind the response
	mSigned := ownrsa.Unblind(sigma, rVal, serverPubK)
	fmt.Print("mSigned: ")
	fmt.Println(mSigned)

	verified := ownrsa.Verify(m, mSigned, serverPubK)
	fmt.Println(verified)

	var iKey int
	for i, k := range keys {
		if k.PubK == packPubK {
			iKey = i
			//save to k the key updated
			k.PubKSigned = ownrsa.ArrayIntToString(mSigned, "_")
			k.Verified = verified
		}
		fmt.Println(k)
	}
	keys[iKey].PubKSigned = ownrsa.ArrayIntToString(mSigned, "_")
	keys[iKey].Verified = verified
	fmt.Println(keys)
	saveKeys(keys, "keys.json")

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

func Verify(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	packPubK := vars["pubK"]
	color.Green(packPubK)

	//read the keys stored in /keys directory
	keys := readKeys("keys.json")

	var key ownrsa.PackRSA
	//search for complete key
	for _, k := range keys {
		if k.PubK == packPubK {
			key = k
		}
	}

	//get the serverIDsigner pubK
	serverPubK := getServerPubK(config.Server)
	m := ownrsa.StringToArrayInt(key.PubK, "_")
	mSigned := ownrsa.StringToArrayInt(key.PubKSigned, "_")

	verified := ownrsa.Verify(m, mSigned, serverPubK)
	fmt.Println(verified)

	for _, k := range keys {
		if k.PubK == packPubK {
			//save to k the key updated
			k.PubKSigned = ownrsa.ArrayIntToString(mSigned, "_")
			k.Verified = verified
		}
	}
	saveKeys(keys, "keys.json")

	jResp, err := json.Marshal(keys)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
