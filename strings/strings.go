package strings

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GenerateRandom(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(b)
}

func IsAlphaNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func IsNumeric(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func IsAlphabet(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}

func FilterAlphabet(s string) string {
	var result strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func FilterNumeric(s string) string {
	var result strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func FilterAlphaNumeric(s string) string {
	var result strings.Builder
	for _, char := range s {
		if unicode.IsLetter(char) && unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}
