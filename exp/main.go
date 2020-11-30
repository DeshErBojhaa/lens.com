package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	hashByte := []byte("this is the string to hash")
	h := hmac.New(sha256.New, []byte("secret"))
	h.Write(hashByte)

	b := h.Sum(nil)

	fmt.Println(b)
}
