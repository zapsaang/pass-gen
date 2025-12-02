package passgen

import (
	"crypto/rand"
	"errors"
	"math/big"
)

const (
	DefaultSaltLength = 32
	randomCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateRandomString(length int) (string, error) {
	if length <= 0 {
		length = DefaultSaltLength
	}

	if length > 4096 {
		return "", errors.New("random string length too large (max 4096)")
	}

	b := make([]byte, length)
	charsetLen := big.NewInt(int64(len(randomCharset)))

	for i := range b {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		b[i] = randomCharset[num.Int64()]
	}

	return string(b), nil
}