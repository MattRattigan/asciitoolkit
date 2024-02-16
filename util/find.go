package util

import "regexp"

func FindASCIIDigits() {

}

// FindASCIIWords finds and returns all ASCII word sequences in a given string.
// A word is defined as a sequence of one or more ASCII alphabetic characters (letters).
// This definition can be adjusted as needed by modifying the regular expression.
func FindASCIIWords(s string) []string {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	return re.FindAllString(s, -1)
}
