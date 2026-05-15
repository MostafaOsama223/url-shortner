package service

import (
	"fmt"
	"math"
)

const SHORT_URL_LENGTH = 62
const CHARSET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base16ToBase10(hex string) int64 {
	var result int64
	len := len(hex)
	for i, c := range hex {
		if c >= '0' && c <= '9' {
			result += (int64(c) - 48) * int64(math.Pow(16, float64(len-i-1)))
		} else if (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
			result += (int64(c) - 87) * int64(math.Pow(16, float64(len-i-1)))
		} else {
			panic(fmt.Sprintf("Invalid hex character: %c", c))
		}
	}

	return result
}

func Base10ToBase62(num int64) string {
	var result []byte

	for num > 0 {
		quotient := num / SHORT_URL_LENGTH
		remainder := num % SHORT_URL_LENGTH
		result = append([]byte{CHARSET[remainder]}, result...)
		num = quotient
	}

	return string(result)
}
