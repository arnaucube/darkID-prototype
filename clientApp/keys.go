package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ownrsa "./ownrsa"
	"github.com/fatih/color"
)

func readKeys(path string) []ownrsa.PackRSA {
	var keys []ownrsa.PackRSA

	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &keys)

	return keys
}

func saveKeys(keys []ownrsa.PackRSA, path string) {
	jsonKeys, err := json.Marshal(keys)
	check(err)
	err = ioutil.WriteFile(path, jsonKeys, 0644)
	check(err)
}

func getServerPubK(url string) ownrsa.RSAPublicKey {
	r, err := http.Get(url + "/")
	check(err)
	fmt.Println(r)

	decoder := json.NewDecoder(r.Body)
	//var sigmaString string
	var pubK ownrsa.RSAPublicKey
	err = decoder.Decode(&pubK)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	color.Blue("received server pubK:")
	fmt.Println(pubK)
	return pubK
}
