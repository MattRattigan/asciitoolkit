package util

import (
	"strings"
	"unicode"
)

// FilterASCIIDigits returns a string containing only the ASCII digit characters from the input string.
func FilterASCIIDigits(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) && r <= '9' {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// FilterASCIILetters returns a string containing only the ASCII letter characters from the input string.
func FilterASCIILetters(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) && ((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// FilterASCIIUpper returns a string containing only the ASCII uppercase letter characters from the input string.
func FilterASCIIUpper(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsUpper(r) && r <= 'Z' {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// FilterASCIILower returns a string containing only the ASCII lowercase letter characters from the input string.
func FilterASCIILower(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLower(r) && r >= 'a' {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// FilterASCIIPunct returns a string containing only the ASCII punctuation characters from the input string.
func FilterASCIIPunct(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsPunct(r) && ((r >= '!' && r <= '/') || (r >= ':' && r <= '@') || (r >= '[' && r <= '`') || (r >= '{' && r <= '~')) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
