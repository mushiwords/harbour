package random

import (
	"crypto/rand"
	"math/big"
	mathrandom "math/rand"
	"time"
)

var (
	character = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	chLen     = len(character)
)

func GetRandomString(size int) string {
	buf := make([]byte, size, size)
	max := big.NewInt(int64(chLen))
	for i := 0; i < size; i++ {
		random, err := rand.Int(rand.Reader, max)
		if err != nil {
			mathrandom.Seed(time.Now().UnixNano())
			buf[i] = character[mathrandom.Intn(chLen)]
			continue
		}
		buf[i] = character[random.Int64()]
	}
	return string(buf)
}

func GetRandomBytes(size int) []byte {
	buf := make([]byte, size, size)

	if _, err := rand.Read(buf); err != nil {
		mathrandom.Seed(time.Now().UnixNano())
		mathrandom.Read(buf)
	}
	return buf
}
