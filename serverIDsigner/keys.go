package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"os"
	"time"
)

type Key struct {
	ID       string    `json:"id"`
	PrivK    string    `json:"privK"` //path of the PrivK file
	PubK     string    `json:"pubK"`  //path of the PubK file
	Date     time.Time `json:"date"`
	Verified bool      `json:"verified"`
	Signed   string    `json:"signed"`
}

func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	check(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	check(err)
}
func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	check(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	check(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	check(err)
}
func openPEMKey(path string) (key rsa.PrivateKey) {
	return
}
func openPublicPEMKey(path string) (key rsa.PublicKey) {
	return
}
func readKeys() []Key {
	path := keysDir + "/keys.json"
	var keys []Key

	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &keys)

	return keys
}

func saveKeys(keys []Key) {
	jsonKeys, err := json.Marshal(keys)
	check(err)
	err = ioutil.WriteFile(keysDir+"/keys.json", jsonKeys, 0644)
	check(err)
}
func getKeyByKeyID(keyID string) (k Key) {
	keys := readKeys()
	for _, key := range keys {
		if key.ID == keyID {
			k = key
		}
	}
	return k
}

/*
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
*/
