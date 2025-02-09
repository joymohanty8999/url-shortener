package utils

import (
	"crypto/rand" // better than math/rand for generating random numbers as it uses system-generated entropy rather than a fixed seed
	"math/big"
	"strings"
)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomBase62String(length int) (string, error) {
	var sb strings.Builder
	base := big.NewInt(int64(len(base62Chars)))

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, base)
		if err != nil {
			return "", err
		}
		sb.WriteByte(base62Chars[num.Int64()])
	}

	return sb.String(), nil
}
