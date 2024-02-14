package letters

import (
	"strings"
)

// Digits provides 0-9 ASCII digits as a string.
func Digits() string {
	return "0123456789"
}

// ASCIILetters generates ASCII characters between ('a' through 'z') and ('A' through 'Z') than returns string.
func ASCIILetters() string {
	var builder strings.Builder
	for lc := 'a'; lc <= 'z'; lc++ {
		builder.WriteRune(lc)
	}
	for uc := 'A'; uc <= 'Z'; uc++ {
		builder.WriteRune(uc)
	}
	return builder.String()
}

// ASCIILowercase generates ASCII lowercase letters that are between ('a' through 'z') than returns string.
func ASCIILowercase() string {
	var builder strings.Builder
	for c := 'a'; c <= 'z'; c++ {
		builder.WriteByte(byte(c))
	}
	return builder.String()
}

// ASCIIUppercase generates ASCII uppercase letters ('A' through 'Z') than returns string.
func ASCIIUppercase() string {
	var builder strings.Builder
	for c := 'A'; c <= 'Z'; c++ {
		builder.WriteByte(byte(c))
	}
	return builder.String()
}

// ASCIIPunctuation generates ASCII punctuation characters that are in the ranges of
// 33-47, 58-64, 91-96, and 123-126 than returns string.
func ASCIIPunctuation() string {
	var builder strings.Builder
	ranges := [][]int{{33, 47}, {58, 64}, {91, 96}, {123, 126}}

	for _, r := range ranges {
		for c := r[0]; c <= r[1]; c++ {
			builder.WriteByte(byte(c))
		}
	}
	return builder.String()
}

// HexDigits generates ASCII characters between ('a' through 'z') and ('A' through 'Z') than returns string
func HexDigits() string {
	var builder strings.Builder
	for d := '0'; d <= '9'; d++ {
		builder.WriteRune(d)
	}
	for lc := 'a'; lc <= 'f'; lc++ {
		builder.WriteRune(lc)
	}
	for uc := 'A'; uc <= 'F'; uc++ {
		builder.WriteRune(uc)
	}
	return builder.String()
}

// OctDigits generates ASCII characters between ('a' through 'z') and ('A' through 'Z') than returns string
func OctDigits() string {
	var builder strings.Builder
	for d := '0'; d <= '7'; d++ {
		builder.WriteRune(d)
	}
	return builder.String()
}

// Printable generates ASCII characters between ('a' through 'z') and ('A' through 'Z') than returns string
func Printable() string {
	var builder strings.Builder
	for d := '0'; d <= '9'; d++ {
		builder.WriteRune(d)
	}
	for uc := 'A'; uc <= 'Z'; uc++ {
		builder.WriteRune(uc)
	}
	for lc := 'a'; lc <= 'z'; lc++ {
		builder.WriteRune(lc)
	}
	punctuationAndWhitespace := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\f\v"
	builder.WriteString(punctuationAndWhitespace)

	return builder.String()
}

// Whitespace generates ASCII characters between ('a' through 'z') and ('A' through 'Z') than returns string
func Whitespace() string {
	var builder strings.Builder
	// Space, tab, newline, carriage return, form feed, vertical tab
	whitespaceChars := " \t\n\r\f\v"
	builder.WriteString(whitespaceChars)

	return builder.String()
}
