package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/mgo.v2/bson"

	ownrsa "./ownrsa"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    string        `json:"token"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	//TODO return the public key, to allow others verifign signed strings by this server

	jResp, err := json.Marshal(serverRSA.PubK)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

func Signup(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Print("user signup: ")
	fmt.Println(user)

	//save the new project to mongodb
	rUser := User{}
	err = userCollection.Find(bson.M{"email": user.Email}).One(&rUser)
	if err != nil {
		//user not exists
		err = userCollection.Insert(user) //TODO find a way to get the object result when inserting in one line, without need of the two mgo petitions
		err = userCollection.Find(bson.M{"email": user.Email}).One(&user)
	} else {
		//user exists
		fmt.Fprintln(w, "User already registered")
		return
	}

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

func Login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	//TODO check if the user password exists in the database

	fmt.Print("user login: ")
	fmt.Println(user)
	token, err := newToken()
	check(err)
	user.Token = token

	//save the new project to mongodb
	rUser := User{}
	err = userCollection.Find(bson.M{"email": user.Email}).One(&rUser)
	if err != nil {
	} else {
		//user exists, update with the token
		err = userCollection.Update(bson.M{"_id": rUser.Id}, user)
		check(err)
	}

	jResp, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

type Sign struct {
	M string `json:"m"`
	C string `json:"c"`
}

type AskBlindSign struct {
	/*PubKString ownrsa.RSAPublicKeyString `json:"pubKstring"`
	PubK       ownrsa.RSAPublicKey       `json:"pubK"`*/
	M string `json:"m"`
}

func BlindSign(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var askBlindSign AskBlindSign
	err := decoder.Decode(&askBlindSign)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	color.Red(askBlindSign.M)
	fmt.Println(askBlindSign)

	/*fmt.Println(askBlindSign)
	askBlindSign.PubK, err = ownrsa.PubKStringToBigInt(askBlindSign.PubKString)
	if err != nil {
		fmt.Fprintln(w, "error")
		return
	}*/

	//convert msg to []int
	/*var m []int
	mBytes := []byte(askBlindSign.M)
	for _, byte := range mBytes {
		m = append(m, int(byte))
	}*/

	m := ownrsa.StringToArrayInt(askBlindSign.M, "_")

	sigma := ownrsa.BlindSign(m, serverRSA.PrivK) //here the privK will be the CA privK, not the m emmiter's one. The pubK is the user's one
	fmt.Print("Sigma': ")
	fmt.Println(sigma)
	sigmaString := ownrsa.ArrayIntToString(sigma, "_")
	askBlindSign.M = sigmaString

	jResp, err := json.Marshal(askBlindSign)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

type PetitionVerifySign struct {
	M       string `json:"m"`
	MSigned string `json:"mSigned"`
}

func VerifySign(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var petitionVerifySign PetitionVerifySign
	err := decoder.Decode(&petitionVerifySign)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Println(petitionVerifySign)

	//convert M to []int
	var mOriginal []int
	mBytes := []byte(petitionVerifySign.M)
	for _, byte := range mBytes {
		mOriginal = append(mOriginal, int(byte))
	}

	//convert MSigned to []int
	var mSignedInts []int
	mSignedString := strings.Split(petitionVerifySign.MSigned, " ")
	for _, s := range mSignedString {
		i, err := strconv.Atoi(s)
		check(err)
		mSignedInts = append(mSignedInts, i)
	}

	verified := ownrsa.Verify(mOriginal, mSignedInts, serverRSA.PubK)

	fmt.Fprintln(w, verified)
}
