package util

import (
	"math/rand"
	"strings"
)

const alphabetSmall = "abcdefghijklmnopqrstuvwxyz"
const alphabetCapital = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabetCapital)
	for i := 0; i < n; i++ {
		if i == 0 {
			c := alphabetCapital[rand.Intn(k)]
			sb.WriteByte(c)
		} else {
			c := alphabetSmall[rand.Intn(k)]
			sb.WriteByte(c)
		}
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomCurrency() string {
	currency := []string{"ETB", "EUR", "USD"}
	return currency[rand.Intn(len(currency))]
}
