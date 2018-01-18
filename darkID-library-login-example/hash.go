package main

import (
	"crypto/sha256"
	"encoding/base64"
)

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
