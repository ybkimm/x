package passwords

import (
	"crypto/hmac"

	"golang.org/x/crypto/sha3"
)

func CreateHash(secret []byte, p []byte) []byte {
	hasher := hmac.New(sha3.New512, secret)
	hasher.Write(p)
	return hasher.Sum(nil)
}
