package hash

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"hash"
)

func NewHMAC(key string) HMAC {
	h := hmac.New(sha512.New, []byte(key))
	return HMAC{
		hmac: h,
	}
}

type HMAC struct {
	hmac hash.Hash
}

func (h HMAC) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	b := h.hmac.Sum(nil)

	return base64.URLEncoding.EncodeToString(b)
}
