package gen

import (
	"math/rand"
	"strings"
	"time"
)

// GenRandomStrRune generate string based on string length
func GenRandomStrRune(length int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var str strings.Builder
	for i := 0; i < length; i++ {
		str.WriteRune(chars[r.Intn(len(chars))])
	}

	return str.String()

}

// GenRandomStrSpec generate string based on string length
// using special characters
func GenRandomStrSpec(length int) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials

	buf := make([]byte, length)
	buf[0] = digits[r.Intn(len(digits))]
	buf[1] = specials[r.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[r.Intn(len(all))]
	}
	r.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)

}
