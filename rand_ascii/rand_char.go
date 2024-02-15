package rand_ascii

import (
	"github.com/matt/aciitools/ascii"
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
// defined in 'ascii.Digits'
func RandCharDigit() string {
	return helperChar(ascii.Digits)
}

// RandCharLetter generates a random ASCII character returns as type string. The selection is made from the set of characters
// defined in 'ascii.ASCIILetters'
func RandCharLetter() string {
	return helperChar(ascii.ASCIILetters)
}

// RandCharLower generates a random ASCII Uppercase character, returns as type string. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIILowercase'
func RandCharLower() string {
	return helperChar(ascii.ASCIILowercase)
}

// RandCharUpper generates a random ASCII Uppercase character returns as type string. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIIUppercase'
func RandCharUpper() string {
	return helperChar(ascii.ASCIIUppercase)
}

// RandCharPunct generates a random ASCII punctuation character returns as type string. The selection is made from the set of punctuation characters
// defined in 'ascii.ASCIIPunctuation'.
func RandCharPunct() string {
	return helperChar(ascii.ASCIIPunctuation)
}
