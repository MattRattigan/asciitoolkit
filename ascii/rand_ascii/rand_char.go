package rand_ascii

import (
	"github.com/MattRattigan/asciitoolkit/ascii"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// helperChar facilitates the actions performed in RandCharDigitStr, RandASCIILetterStr, RandASCIILowerStr, RandASCIIUpperStr,
// and RandASCIIPunctStr
func helperChar(f func() string) string {
	val := f()
	return string(val[rand.Intn(len(val))])
}

// RandCharDigitStr generates a random ASCII character returns as type string. The selection is made from the set of characters
// defined in 'ascii.Digits'
func RandCharDigitStr() string {
	return helperChar(ascii.Digits)
}

// RandASCIILetterStr generates a random ASCII character returns as type string. The selection is made from the set of characters
// defined in 'ascii.ASCIILetters'
func RandASCIILetterStr() string {
	return helperChar(ascii.ASCIILetters)
}

// RandASCIILowerStr generates a random ASCII Uppercase character, returns as type string. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIILowercase'
func RandASCIILowerStr() string {
	return helperChar(ascii.ASCIILowercase)
}

// RandASCIIUpperStr generates a random ASCII Uppercase character returns as type string. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIIUppercase'
func RandASCIIUpperStr() string {
	return helperChar(ascii.ASCIIUppercase)
}

// RandASCIIPunctStr generates a random ASCII punctuation character returns as type string. The selection is made from the set of punctuation characters
// defined in 'ascii.ASCIIPunctuation'.
func RandASCIIPunctStr() string {
	return helperChar(ascii.ASCIIPunctuation)
}
