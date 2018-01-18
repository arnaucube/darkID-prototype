package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
)

type Key struct {
	ID             string         `json:"id"`
	PrivK          string         `json:"privK"` //path of the PrivK file
	PubK           string         `json:"pubK"`  //path of the PubK file
	PublicKey      string         `json:"publicKey"`
	Date           time.Time      `json:"date"`
	Hashed         []byte         `json:"hashed"`
	UnblindedSig   []byte         `json:"unblindedsig"`
	Verified       bool           `json:"verified"`
	ServerVerifier *rsa.PublicKey `json:"serververifier"`
	SignerID       string         `json:"signerid"`
	BlockchainRef  string         `json:"blockchainref"`
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey rsa.PublicKey) (string, error) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	check(err)
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: asn1Bytes,
		},
	)
	color.Red("pubkey_pem")
	fmt.Println(pubkey_pem)
	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (pub rsa.PublicKey, err error) {
	pemBlock, _ := pem.Decode([]byte(pubPEM))
	_, err = asn1.Unmarshal(pemBlock.Bytes, &pub)
	return
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
func openPEMKey(path string) (key *rsa.PrivateKey, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	key, err = ParseRsaPrivateKeyFromPemStr(string(b))
	return
}
func openPublicPEMKey(path string) (key rsa.PublicKey, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	key, err = ParseRsaPublicKeyFromPemStr(string(b))
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
func saveKey(k Key) {
	fmt.Println(k)
	keys := readKeys()
	for i, key := range keys {
		if key.ID == k.ID {
			keys[i] = k
		}
	}
	saveKeys(keys)
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
func removeKey(keyID string, originalKeys []Key) (keys []Key) {
	for _, key := range originalKeys {
		if key.ID != keyID {
			keys = append(keys, key)
		}
	}
	return
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
