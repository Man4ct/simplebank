package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomString(n int) string {
	var stringBuilder strings.Builder

	alphabetLength := len(alphabet)
	for i := 0; i < n; i++ {
		randAlphabet := alphabet[rand.Intn(alphabetLength)]

		stringBuilder.WriteByte(randAlphabet)
	}

	return stringBuilder.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]

}
