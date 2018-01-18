package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "this is the secret signing key"
)

var createdToken string

func initializeToken() {
	var err error
	createdToken, err = newToken()
	if err != nil {
		fmt.Println("Creating token failed")
	}
}

func newToken() (string, error) {
	signingKeyB := []byte(signingKey)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(signingKeyB)
	return tokenString, err
}

func parseToken(myToken string, myKey string) {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(myKey), nil
	})

	if err == nil && token.Valid {
		fmt.Println("Your token is valid.  I like your style.")
	} else {
		fmt.Println("This token is terrible!  I cannot accept this.")
	}
}
