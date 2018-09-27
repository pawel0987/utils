// Author: Paweł Konopko
// License: MIT

package random_utils

import (
	"math/rand"
)

func SetSeed(seed int64) {
	rand.Seed(seed)
}

func GenerateRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*(){}[];:'.,/?+=-_0987654321"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
