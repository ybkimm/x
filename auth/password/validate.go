package password

import (
	"crypto/hmac"

	"golang.org/x/crypto/sha3"
)

func Validate(secret []byte, h, p []byte) bool {
	hasher := hmac.New(sha3.New512, secret)
	hasher.Write(p)
	expected := hasher.Sum(nil)
	return hmac.Equal(expected, h)
}
