package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
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

	time.Sleep(time.Second * 2)

	b, err := ioutil.ReadFile(keysDir + "/" + key.PubK)
	if err != nil {
		fmt.Print(err)
	}
	key.PublicKey = string(b)

	key.Date = time.Now()
	fmt.Println(key.PublicKey)

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

func Delete(keyID string) []Key {
	originalKeys := readKeys()
	//remove key .pem files
	key := getKeyByKeyID(keyID)
	_, err := exec.Command("rm", keysDir+"/"+key.PrivK).CombinedOutput()
	check(err)
	_, err = exec.Command("rm", keysDir+"/"+key.PubK).CombinedOutput()
	check(err)

	//remove key from keys.json
	keys := removeKey(keyID, originalKeys)
	saveKeys(keys)
	return keys
}

func Encrypt(keyID string, encryptData EncryptData) EncryptData {
	key := getKeyByKeyID(keyID)
	pubK, err := openPublicPEMKey(keysDir + "/" + key.PubK)
	check(err)
	out, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, &pubK, []byte(encryptData.M), []byte("orders"))
	check(err)
	fmt.Println(string(out))
	encryptData.C = out
	return encryptData
}

func Decrypt(keyID string, encryptData EncryptData) EncryptData {
	key := getKeyByKeyID(keyID)
	privK, err := openPEMKey(keysDir + "/" + key.PrivK)
	check(err)
	out, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, privK, []byte(encryptData.C), []byte("orders"))
	check(err)
	fmt.Println(string(out))
	encryptData.M = string(out)
	return encryptData
}
