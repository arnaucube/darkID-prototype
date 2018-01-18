package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	mrand "math/rand"
	"net/http"
	"strings"

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

type Proof struct {
	PublicKey string `json:"publicKey"`
	Clear     string `json:"clear"`
	Question  []byte `json:"question"`
	Answer    string `json:"answer"`
}

var proofs []Proof

func GetProof(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var receivedProof Proof
	err := decoder.Decode(&receivedProof)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//TODO check if the user password exists in the database

	stringPublicKey := strings.Replace(receivedProof.PublicKey, " ", "\n", -1)
	stringPublicKey = strings.Replace(stringPublicKey, "-----BEGIN\n", "-----BEGIN ", -1)
	stringPublicKey = strings.Replace(stringPublicKey, "-----END\n", "-----END ", -1)
	stringPublicKey = strings.Replace(stringPublicKey, "PUBLIC\n", "PUBLIC ", -1)
	color.Green(stringPublicKey)
	publicKey, err := ParseRsaPublicKeyFromPemStr(stringPublicKey)
	check(err)

	var proof Proof
	proof.Clear = RandStringRunes(40)

	out, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, &publicKey, []byte(proof.Clear), []byte("orders"))
	check(err)
	proof.Question = out

	proofs = append(proofs, proof)

	proof.Clear = ""
	jResp, err := json.Marshal(proof)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}
func AnswerProof(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var ansProof Proof
	err := decoder.Decode(&ansProof)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	proof, err := getProofFromStorage(ansProof.PublicKey)
	if err != nil {

	}
	if ansProof.Answer == proof.Clear {
		token, err := newToken()
		check(err)
		fmt.Fprintln(w, string(token))
	}

	fmt.Fprintln(w, string("fail"))
}
func getProofFromStorage(publicKey string) (Proof, error) {
	var voidProof Proof
	for _, proof := range proofs {
		if proof.PublicKey == publicKey {
			return proof, nil
		}
	}
	return voidProof, errors.New("proof not exist in storage")
}

//function to generate random string of fixed length
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[mrand.Intn(len(letterRunes))]
	}
	return string(b)
}
