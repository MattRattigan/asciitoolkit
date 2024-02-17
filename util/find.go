package util

import "regexp"

// FindASCIIDigits finds and returns all sequences of ASCII digits in the given string.
func FindASCIIDigits(s string) []string {
	digitRegexp := regexp.MustCompile(`\d+`)
	return digitRegexp.FindAllString(s, -1)
}

// FindASCIIWords finds and returns all ASCII word sequences in a given string.
// A word is defined as a sequence of one or more ASCII alphabetic characters (letters).
// This definition can be adjusted as needed by modifying the regular expression.
func FindASCIIWords(s string) []string {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	return re.FindAllString(s, -1)
}
