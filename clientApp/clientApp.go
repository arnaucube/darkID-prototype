package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cryptoballot/fdh"
	"github.com/cryptoballot/rsablind"
)

func IDs() []Key {
	//read the keys stored in /keys directory
	keys := readKeys()
	return keys
}
func NewID() []Key {
	//generate RSA keys pair
	reader := rand.Reader
	k, err := rsa.GenerateKey(reader, keysize)
	check(err)

	id := hash(time.Now().String())
	savePEMKey(keysDir+"/"+id+"private.pem", k)
	savePublicPEMKey(keysDir+"/"+id+"public.pem", k.PublicKey)

	var key Key
	key.ID = id
	key.PrivK = id + "private.pem"
	key.PubK = id + "public.pem"

	key.Date = time.Now()
	fmt.Println(key)

	keys := readKeys()
	keys = append(keys, key)
	saveKeys(keys)
	return keys
}

type AskBlindSign struct {
	M []byte `json:"m"`
}
type SignResponse struct {
	Sig  []byte        `json:"sig"`
	PubK rsa.PublicKey `json:"pubK"`
}

func BlindAndSendToSign(keyID string) []byte {
	//get the key
	key := getKeyByKeyID(keyID)
	//privK := openPEMKey(key.PrivK)
	pubK := openPublicPEMKey(key.PubK)

	//TODO pubK to string
	m := []byte("pubK") //convert pubK to array of bytes
	//blind the hashed message
	// We do a SHA256 full-domain-hash expanded to 1536 bits (3/4 the key size)
	hashed := fdh.Sum(crypto.SHA256, hashize, m)
	blinded, unblinder, err := rsablind.Blind(&pubK, hashed)
	if err != nil {
		panic(err)
	}
	var askBlindSign AskBlindSign
	askBlindSign.M = blinded
	//send blinded to serverIDsigner
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(askBlindSign)
	res, err := http.Post(config.Server+"blindsign", "application/json", body)
	check(err)
	var signResponse SignResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&signResponse)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	sig := signResponse.Sig
	serverPubK := signResponse.PubK

	//unblind the signedblind
	unblindedSig := rsablind.Unblind(&serverPubK, sig, unblinder)
	fmt.Println(unblindedSig)

	return unblindedSig
}

func Verify(packPubK string) {

	return
}
