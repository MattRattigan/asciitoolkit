package rand_ascii

import (
	"github.com/MattRattigan/asciitoolkit/ascii"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Iota
const (
	// LOW flag representation of lowercase ASCII
	LOW = 1 << iota
	// UPP flag representation of uppercase ASCII
	UPP
	// DIG flag representation of 0-9
	DIG
	// PUNCT flag representation or punctuation ASCII characters
	PUNCT
)

// helper facilitates the actions performed in RandASCIIDigit, RandASCIILetter, RandASCIILower, RandASCIIUpper,
// and RandASCIIPunct
func helper(f func() string) rune {
	val := f()
	return rune(val[rand.Intn(len(val))])
}

// RandASCIIDigit generates a random ASCII character returns as type rune. The selection is made from the set of characters
// defined in 'ascii.Digits'
func RandASCIIDigit() rune {
	return helper(ascii.Digits)
}

// RandASCIILetter generates a random ASCII character returns as type rune. The selection is made from the set of characters
// defined in 'ascii.ASCIILetters'
func RandASCIILetter() rune {
	return helper(ascii.ASCIILetters)
}

// RandASCIILower generates a random ASCII Uppercase character returns as type rune. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIILowercase'
func RandASCIILower() rune {
	return helper(ascii.ASCIILowercase)
}

// RandASCIIUpper generates a random ASCII Uppercase character returns as type rune. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIIUppercase'
func RandASCIIUpper() rune {
	return helper(ascii.ASCIIUppercase)
}

// RandASCIIPunct generates a random ASCII punctuation character returns as type rune. The selection is made from the set of punctuation characters
// defined in 'ascii.ASCIIPunctuation'.
func RandASCIIPunct() rune {
	return helper(ascii.ASCIIPunctuation)
}

// RandASCIIString generates a random ASCII string of the specified length. The 'flags' parameter
// determines which categories of ASCII characters (lowercase, uppercase, digits, punctuation)
// can be included in the result. Multiple categories can be combined using bitwise OR ('|').
func RandASCIIString(length int, flags int) string {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		switch rand.Intn(4) {
		case 0:
			if flags&DIG != 0 {
				builder.WriteRune(RandASCIIDigit())
				continue
			}
		case 1:
			if flags&LOW != 0 {
				builder.WriteRune(RandASCIILower())
				continue
			}
		case 2:
			if flags&UPP != 0 {
				builder.WriteRune(RandASCIIUpper())
				continue
			}
		case 3:
			if flags&PUNCT != 0 {
				builder.WriteRune(RandASCIIPunct())
				continue
			}
		}

		if flags&DIG != 0 {
			builder.WriteRune(RandASCIIDigit())
		} else if flags&LOW != 0 {
			builder.WriteRune(RandASCIILower())
		} else if flags&UPP != 0 {
			builder.WriteRune(RandASCIIUpper())
		} else if flags&PUNCT != 0 {
			builder.WriteRune(RandASCIIPunct())
		}
	}

	return builder.String()
}
