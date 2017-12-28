package ownrsa

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type RSAPublicKey struct {
	E *big.Int `json:"e"`
	N *big.Int `json:"n"`
}
type RSAPublicKeyString struct {
	E string `json:"e"`
	N string `json:"n"`
}
type RSAPrivateKey struct {
	D *big.Int `json:"d"`
	N *big.Int `json:"n"`
}

type RSA struct {
	PubK  RSAPublicKey
	PrivK RSAPrivateKey
}

type PackRSA struct {
	PubK       string    `json:"pubK"`
	PrivK      string    `json:"privK"`
	Date       time.Time `json:"date"`
	PubKSigned string    `json:"pubKSigned"`
}

const maxPrime = 500
const minPrime = 100

func GenerateKeyPair() RSA {

	rand.Seed(time.Now().Unix())
	p := randPrime(minPrime, maxPrime)
	q := randPrime(minPrime, maxPrime)
	fmt.Print("p:")
	fmt.Println(p)
	fmt.Print("q:")
	fmt.Println(q)

	n := p * q
	phi := (p - 1) * (q - 1)
	e := 65537
	var pubK RSAPublicKey
	pubK.E = big.NewInt(int64(e))
	pubK.N = big.NewInt(int64(n))

	d := new(big.Int).ModInverse(big.NewInt(int64(e)), big.NewInt(int64(phi)))

	var privK RSAPrivateKey
	privK.D = d
	privK.N = big.NewInt(int64(n))

	var rsa RSA
	rsa.PubK = pubK
	rsa.PrivK = privK
	return rsa
}
func Encrypt(m string, pubK RSAPublicKey) []int {
	var c []int
	mBytes := []byte(m)
	for _, byte := range mBytes {
		c = append(c, EncryptInt(int(byte), pubK))
	}
	return c
}
func Decrypt(c []int, privK RSAPrivateKey) string {
	var m string
	var mBytes []byte
	for _, indC := range c {
		mBytes = append(mBytes, byte(DecryptInt(indC, privK)))
	}
	m = string(mBytes)
	return m
}

func EncryptBigInt(bigint *big.Int, pubK RSAPublicKey) *big.Int {
	Me := new(big.Int).Exp(bigint, pubK.E, nil)
	c := new(big.Int).Mod(Me, pubK.N)
	return c
}
func DecryptBigInt(bigint *big.Int, privK RSAPrivateKey) *big.Int {
	Cd := new(big.Int).Exp(bigint, privK.D, nil)
	m := new(big.Int).Mod(Cd, privK.N)
	return m
}

func EncryptInt(char int, pubK RSAPublicKey) int {
	charBig := big.NewInt(int64(char))
	Me := charBig.Exp(charBig, pubK.E, nil)
	c := Me.Mod(Me, pubK.N)
	return int(c.Int64())
}
func DecryptInt(val int, privK RSAPrivateKey) int {
	valBig := big.NewInt(int64(val))
	Cd := valBig.Exp(valBig, privK.D, nil)
	m := Cd.Mod(Cd, privK.N)
	return int(m.Int64())
}

func Blind(m []int, r int, pubK RSAPublicKey, privK RSAPrivateKey) []int {
	var mBlinded []int
	rBigInt := big.NewInt(int64(r))
	for i := 0; i < len(m); i++ {
		mBigInt := big.NewInt(int64(m[i]))
		rE := new(big.Int).Exp(rBigInt, pubK.E, nil)
		mrE := new(big.Int).Mul(mBigInt, rE)
		mrEmodN := new(big.Int).Mod(mrE, privK.N)
		mBlinded = append(mBlinded, int(mrEmodN.Int64()))
	}
	return mBlinded
}

func BlindSign(m []int, privK RSAPrivateKey) []int {
	var r []int
	for i := 0; i < len(m); i++ {
		mBigInt := big.NewInt(int64(m[i]))
		sigma := new(big.Int).Exp(mBigInt, privK.D, privK.N)
		r = append(r, int(sigma.Int64()))
	}
	return r
}
func Unblind(blindsigned []int, r int, pubK RSAPublicKey) []int {
	var mSigned []int
	rBigInt := big.NewInt(int64(r))
	for i := 0; i < len(blindsigned); i++ {
		bsBigInt := big.NewInt(int64(blindsigned[i]))
		//r1 := new(big.Int).Exp(rBigInt, big.NewInt(int64(-1)), nil)
		r1 := new(big.Int).ModInverse(rBigInt, pubK.N)
		bsr := new(big.Int).Mul(bsBigInt, r1)
		sig := new(big.Int).Mod(bsr, pubK.N)
		mSigned = append(mSigned, int(sig.Int64()))
	}
	return mSigned
}
func Verify(msg []int, mSigned []int, pubK RSAPublicKey) bool {
	if len(msg) != len(mSigned) {
		return false
	}
	var mSignedDecrypted []int
	for _, ms := range mSigned {
		msBig := big.NewInt(int64(ms))
		//decrypt the mSigned with pubK
		Cd := new(big.Int).Exp(msBig, pubK.E, nil)
		m := new(big.Int).Mod(Cd, pubK.N)
		mSignedDecrypted = append(mSignedDecrypted, int(m.Int64()))
	}
	fmt.Print("msg signed decrypted: ")
	fmt.Println(mSignedDecrypted)
	r := true
	//check if the mSignedDecrypted == msg
	for i := 0; i < len(msg); i++ {
		if msg[i] != mSignedDecrypted[i] {
			r = false
		}
	}
	return r
}

func HomomorphicMultiplication(c1 int, c2 int, pubK RSAPublicKey) int {
	c1BigInt := big.NewInt(int64(c1))
	c2BigInt := big.NewInt(int64(c2))
	c1c2 := new(big.Int).Mul(c1BigInt, c2BigInt)
	n2 := new(big.Int).Mul(pubK.N, pubK.N)
	d := new(big.Int).Mod(c1c2, n2)
	r := int(d.Int64())
	return r
}

func PubKStringToBigInt(kS RSAPublicKeyString) (RSAPublicKey, error) {
	var k RSAPublicKey
	var ok bool
	k.E, ok = new(big.Int).SetString(kS.E, 10)
	if !ok {
		return k, errors.New("error parsing big int E")
	}
	k.N, ok = new(big.Int).SetString(kS.N, 10)
	if !ok {
		return k, errors.New("error parsing big int N")
	}
	return k, nil
}

func PackKey(k RSA) PackRSA {
	var p PackRSA
	p.PubK = k.PubK.E.String() + "," + k.PubK.N.String()
	p.PrivK = k.PrivK.D.String() + "," + k.PrivK.N.String()
	return p
}

func UnpackKey(p PackRSA) RSA {
	var k RSA
	var ok bool
	k.PubK.E, ok = new(big.Int).SetString(strings.Split(p.PubK, ",")[0], 10)
	k.PubK.N, ok = new(big.Int).SetString(strings.Split(p.PubK, ",")[1], 10)
	k.PrivK.D, ok = new(big.Int).SetString(strings.Split(p.PrivK, ",")[0], 10)
	k.PrivK.N, ok = new(big.Int).SetString(strings.Split(p.PrivK, ",")[1], 10)
	if !ok {
		fmt.Println("error on Unpacking Keys")
	}
	return k
}

func ArrayIntToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
func StringToArrayInt(s string, delim string) []int {
	var a []int
	arrayString := strings.Split(s, delim)
	for _, s := range arrayString {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		a = append(a, i)
	}
	return a
}
