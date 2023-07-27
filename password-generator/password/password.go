package password

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func randomChar(chars string) byte {
	return chars[rand.Intn(len(chars))]
}

func GeneratePassword(lenght int) string {
	rand.Seed(time.Now().UnixNano())
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	numbers := "0123456789"
	specials := "@!$%?"

	var password strings.Builder
	password.Grow(lenght)
	for i := 0; i < lenght; i++ {
		if i%4 == 0 {
			password.WriteByte(randomChar(uppercase))
		} else if i%4 == 1 {
			password.WriteByte(randomChar(lowercase))
		} else if i%4 == 2 {
			password.WriteByte(randomChar(numbers))
		} else {
			password.WriteByte(randomChar(specials))
		}
	}

	return password.String()
}

func GetMd5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}
