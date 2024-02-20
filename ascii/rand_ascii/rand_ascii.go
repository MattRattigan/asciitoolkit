package rand_ascii

import (
	"github.com/MattRattigan/asciitoolkit/ascii"
	"math/rand"
	"strings"
	"time"
)

// rand.Seed depreciated Go 1.20 the new way of doing it. A pointer, source and new rand object
var ran *rand.Rand

func init() {
	src := rand.NewSource(time.Now().UnixNano())
	ran = rand.New(src)
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

// helperRune facilitates the actions performed in RandASCIIDigitRune, RandASCIILetterRune, RandASCIILowerRune, RandASCIIUpperRune,
// and RandASCIIPunctRune
func helperRune(f func() string) rune {
	val := f()
	return rune(val[ran.Intn(len(val))])
}

// RandASCIIDigitRune generates a random ASCII character returns as type rune. The selection is made from the set of characters
// defined in 'ascii.Digits'
func RandASCIIDigitRune() rune {
	return helperRune(ascii.Digits)
}

// RandASCIILetterRune generates a random ASCII character returns as type rune. The selection is made from the set of characters
// defined in 'ascii.ASCIILetters'
func RandASCIILetterRune() rune {
	return helperRune(ascii.ASCIILetters)
}

// RandASCIILowerRune generates a random ASCII Uppercase character returns as type rune. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIILowercase'
func RandASCIILowerRune() rune {
	return helperRune(ascii.ASCIILowercase)
}

// RandASCIIUpperRune generates a random ASCII Uppercase character returns as type rune. The selection is made from the set of uppercase characters
// defined in 'ascii.ASCIIUppercase'
func RandASCIIUpperRune() rune {
	return helperRune(ascii.ASCIIUppercase)
}

// RandASCIIPunctRune generates a random ASCII punctuation character returns as type rune. The selection is made from the set of punctuation characters
// defined in 'ascii.ASCIIPunctuation'.
func RandASCIIPunctRune() rune {
	return helperRune(ascii.ASCIIPunctuation)
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
				builder.WriteRune(RandASCIIDigitRune())
				continue
			}
		case 1:
			if flags&LOW != 0 {
				builder.WriteRune(RandASCIILowerRune())
				continue
			}
		case 2:
			if flags&UPP != 0 {
				builder.WriteRune(RandASCIIUpperRune())
				continue
			}
		case 3:
			if flags&PUNCT != 0 {
				builder.WriteRune(RandASCIIPunctRune())
				continue
			}
		}

		if flags&DIG != 0 {
			builder.WriteRune(RandASCIIDigitRune())
		} else if flags&LOW != 0 {
			builder.WriteRune(RandASCIILowerRune())
		} else if flags&UPP != 0 {
			builder.WriteRune(RandASCIIUpperRune())
		} else if flags&PUNCT != 0 {
			builder.WriteRune(RandASCIIPunctRune())
		}
	}

	return builder.String()
}
