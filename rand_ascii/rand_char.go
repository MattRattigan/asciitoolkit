package rand_ascii

import (
	"github.com/matt/aciitools/letters"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// helperChar facilitates the actions performed in RandCharDigit, RandCharLetter, RandCharLower, RandCharUpper,
// and RandCharPunct
func helperChar(f func() string) string {
	val := f()
	return string(val[rand.Intn(len(val))])
}

// RandCharDigit generates a random ASCII character returns as type string. The selection is made from the set of characters
// defined in 'letters.Digits'
func RandCharDigit() string {
	return helperChar(letters.Digits)
}

// RandCharLetter generates a random ASCII character returns as type string. The selection is made from the set of characters
// defined in 'letters.ASCIILetters'
func RandCharLetter() string {
	return helperChar(letters.ASCIILetters)
}

// RandCharLower generates a random ASCII Uppercase character, returns as type string. The selection is made from the set of uppercase characters
// defined in 'letters.ASCIILowercase'
func RandCharLower() string {
	return helperChar(letters.ASCIILowercase)
}

// RandCharUpper generates a random ASCII Uppercase character returns as type string. The selection is made from the set of uppercase characters
// defined in 'letters.ASCIIUppercase'
func RandCharUpper() string {
	return helperChar(letters.ASCIIUppercase)
}

// RandCharPunct generates a random ASCII punctuation character returns as type string. The selection is made from the set of punctuation characters
// defined in 'letters.ASCIIPunctuation'.
func RandCharPunct() string {
	return helperChar(letters.ASCIIPunctuation)
}
