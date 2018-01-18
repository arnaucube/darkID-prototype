package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cryptoballot/rsablind"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Token    string        `json:"token"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	// return server public key, to allow others verifign signed strings by this server
	jResp, err := json.Marshal(serverKey.PublicKey)
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
		jResp, err := json.Marshal("error login, email not found")
		check(err)
		fmt.Fprintln(w, string(jResp))
		return
	}
	//user exists, check password
	if user.Password != rUser.Password {
		jResp, err := json.Marshal("error login, password not match")
		check(err)
		fmt.Fprintln(w, string(jResp))
		return
	}
	//update with the token
	err = userCollection.Update(bson.M{"_id": rUser.Id}, user)
	check(err)

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
	M []byte `json:"m"`
}
type SignResponse struct {
	Sig  []byte        `json:"sig"`
	PubK rsa.PublicKey `json:"pubK"`
}

func BlindSign(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var askBlindSign AskBlindSign
	err := decoder.Decode(&askBlindSign)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println(askBlindSign)
	blinded := askBlindSign.M

	/*privK := openPEMKey(keysDir + "/server_private.pem")
	pubK := openPublicPEMKey(keysDir + "/server_public.pem")*/
	sig, err := rsablind.BlindSign(serverKey, blinded)
	check(err)
	var signResponse SignResponse
	signResponse.Sig = sig
	signResponse.PubK = serverKey.PublicKey

	jResp, err := json.Marshal(signResponse)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(jResp))
}

//TODO verifysign will not be necessary in this server
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

	//verified := ownrsa.Verify(mOriginal, mSignedInts, serverRSA.PubK)
	verified := false
	fmt.Fprintln(w, verified)
}
