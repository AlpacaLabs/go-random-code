package code

import (
	"math/rand"
	"time"
)

// New returns a 6-character random code.
func New() string {
	return GenerateReasonableCode(6)
}

// GenerateReasonableCode generates a random string consisting of unambiguous capital letters
// or numbers. 1, I, L, O, and 0 are removed to reduce confusion when a user receives and types
// in a code. Vowels are also removed to minimize the odds of generating a controversial word
// of the English language.
func GenerateReasonableCode(desiredLength int) string {
	alphabet := GetReasonableEncodingAlphabet()
	return GenerateRandomString(alphabet, desiredLength)
}

// GetReasonableEncodingAlphabet returns a string of unambiguous capital letters and numbers. 1, I, L, O, and 0
// are removed to reduce confusion when a user attempts to type in their code. Vowels are also removed to minimize
// the odds of generating a controversial word of the English language.
func GetReasonableEncodingAlphabet() string {
	return "23456789BCDFGHJKMNPWRSTVWXYZ"
}

// GenerateRandomString generates a random string based on an encoding alphabet and desired length.
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func GenerateRandomString(alphabet string, desiredLength int) string {
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, desiredLength)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := desiredLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(alphabet) {
			b[i] = alphabet[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
