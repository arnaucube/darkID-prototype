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
	"github.com/fatih/color"
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

func BlindAndSendToSign(keyID string) []Key {
	//get the key
	key := getKeyByKeyID(keyID)
	//privK := openPEMKey(key.PrivK)
	pubK, err := openPublicPEMKey(keysDir + "/" + key.PubK)
	check(err)

	//pubK to string
	m, err := ExportRsaPublicKeyAsPemStr(pubK)
	check(err)
	mB := []byte(m)

	//get serverPubK
	var serverPubK *rsa.PublicKey
	res, err := http.Get(config.Server)
	check(err)
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&serverPubK)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	//blind the hashed message
	// We do a SHA256 full-domain-hash expanded to 1536 bits (3/4 the key size)
	hashed := fdh.Sum(crypto.SHA256, hashize, mB)
	blinded, unblinder, err := rsablind.Blind(serverPubK, hashed)
	if err != nil {
		panic(err)
	}
	var askBlindSign AskBlindSign
	askBlindSign.M = blinded
	//send blinded to serverIDsigner
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(askBlindSign)
	res, err = http.Post(config.Server+"blindsign", "application/json", body)
	check(err)
	var signResponse SignResponse
	decoder = json.NewDecoder(res.Body)
	err = decoder.Decode(&signResponse)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	sig := signResponse.Sig
	//serverPubK := signResponse.PubK

	//unblind the signedblind
	unblindedSig := rsablind.Unblind(serverPubK, sig, unblinder)
	color.Green("unblindedSig")
	fmt.Println(unblindedSig)

	// Verify the original hashed message against the unblinded signature
	if err := rsablind.VerifyBlindSignature(serverPubK, hashed, unblindedSig); err != nil {
		fmt.Println(err)
	} else {
		color.Green("blind signature verified")
		key.Verified = true
	}
	key.UnblindedSig = unblindedSig
	key.Hashed = hashed
	key.ServerVerifier = serverPubK
	saveKey(key)
	keys := readKeys()
	return keys
}

func Verify(packPubK string) {

	return
}
