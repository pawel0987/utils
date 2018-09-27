// Author: Paweł Konopko
// License: MIT

package crypt_utils

import (
	"crypto/sha256"
	"encoding/hex"
	"crypto/sha512"
)

func GenerateSHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateSHA512(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
