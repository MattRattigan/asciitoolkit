package letters

// IsASCIIDigit checks if the given rune r represents an ASCII digit between ('0' through '9').
// returns true if r is within the range of ASCII digits, false if otherwise.
func IsASCIIDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// IsASCIILetter checks if the given rune r represents an ASCII letters that reside between ('A' through 'Z') or .
// ('a' through 'z') returns true if r is within the range of ASCII letters, false if otherwise.
func IsASCIILetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

// IsASCIIUpper checks if the given rune r represents an ASCII uppercase letters ('A' through 'Z').
// returns true if r is within the range of ASCII uppercase letters, false if otherwise.
func IsASCIIUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// IsASCIILower checks if the given rune r represents an ASCII lowercase letter ('a' through 'z').
// returns true if r is within the range of ASCII lowercase letters, false if otherwise.
func IsASCIILower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// IsASCIIPunct checks if the given rune r represents an ASCII punctuation character.
// ASCII punctuation characters are in the ranges of 33-47, 58-64, 91-96, and 123-126.
// It returns true if r is within any of these ranges, indicating it is a punctuation character; false otherwise.
func IsASCIIPunct(r rune) bool {
	return (r >= 33 && r <= 47) || (r >= 58 && r <= 64) ||
		(r >= 91 && r <= 96) || (r >= 123 && r <= 126)
}

// IsASCII checks whether character within the provided string are the standard ASCII characters
// (Printable and Control characters), "" will return true.
func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return true
}

// IsASCIIAlphanumeric checks whether the provided string is set of alphanumeric characters,
// letters or digits of Printable and Control characters
func IsASCIIAlphanumeric(s string) bool {
	for _, r := range s {
		if !IsASCIIDigit(r) && !IsASCIILetter(r) {
			return false
		}
	}

	return true
}
