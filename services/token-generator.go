package services

import (
	"math/rand"
	"time"
)
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateToken(lenght int) string {
	buf := make([]rune, lenght)
	for i := range buf {
		buf[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(buf)
}
