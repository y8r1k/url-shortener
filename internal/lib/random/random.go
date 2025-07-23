package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	allowedCharsInURL := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	res := make([]rune, size)
	for i := 0; i < size; i++ {
		res[i] = allowedCharsInURL[rnd.Intn(len(allowedCharsInURL))]
	}
	return string(res)
}
