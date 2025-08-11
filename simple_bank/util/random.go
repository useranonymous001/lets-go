// this package is used to generate random data for testing purposes

package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

// RandomInt generates a random integer between max and min
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates random Owner name for seeding
func RandomOwner() string {
	return RandomString(7)
}

func RandomBalance() int64 {
	return RandomInt(1000, 1000000)
}

func RandomCurrency() string {
	currencies := []string{
		"USD",
		"NPR",
		"CAD",
		"INR",
		"EUR",
		"RIYAL",
	}

	return currencies[rand.Intn(len(currencies))]
}
