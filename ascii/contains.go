package ascii

// helperRune facilitates the actions performed in ContainsASCIIDigit, ContainsASCIILetter, ContainsASCIIUpper, ContainsASCIILower,
// and ContainsASCIIPunct
func helper(s string, f func(r rune) bool) bool {
	for _, r := range s {
		if f(r) {
			return true
		}
	}
	return false
}

// ContainsASCIIDigit checks if the given string s contains at least one ASCII digit ('0'-'9').
// iterates over each character in the string, applying the IsASCIIDigit
// check to determine if a digit is present. Returns true if any ASCII digit is found, false otherwise.
func ContainsASCIIDigit(s string) bool {
	return helper(s, isRuneASCIIDigit)
}

// ContainsASCIILetter checks if the given string s contains at least one ASCII
// letter ('a' through 'z') and ('A' through 'Z'). iterates over each character in the string, applying the
// IsASCIILetter check to determine if a digit is present. Returns true if any ASCII letter is found, false otherwise.
func ContainsASCIILetter(s string) bool {
	return helper(s, isRuneASCIILetter)
}

// ContainsASCIIUpper checks if the given string s contains at least one ASCII
// ('A' through 'Z'). iterates over each character in the string, applying the
// IsASCIIUpper check to determine if uppercase letter is present. Returns true
// if any ASCII letter is found, false otherwise.
func ContainsASCIIUpper(s string) bool {
	return helper(s, isRuneASCIIUpper)
}

// ContainsASCIILower checks if the given string s contains at least one ASCII
// ('a' through 'z'). iterates over each character in the string, applying the
// IsASCIILower check to determine if a lowercase letter is present. Returns true
// if any ASCII letter is found, false otherwise.
func ContainsASCIILower(s string) bool {
	return helper(s, isRuneASCIILower)
}

// ContainsASCIIPunct checks if the given string s contains at least one ASCII
// ('a' through 'z'). iterates over each character in the string, applying the
// IsASCIIPunct check to determine if a lowercase letter is present. Returns true
// if any ASCII letter is found, false otherwise.
func ContainsASCIIPunct(s string) bool {
	return helper(s, isRuneASCIIPunct)
}

func isRuneASCIIDigit(r rune) bool {
	return IsASCIIDigit(r)
}

func isRuneASCIILetter(r rune) bool {
	return IsASCIILetter(r)
}

func isRuneASCIILower(r rune) bool {
	return IsASCIILower(r)
}

func isRuneASCIIUpper(r rune) bool {
	return IsASCIIUpper(r)
}

func isRuneASCIIPunct(r rune) bool {
	return IsASCIIPunct(r)
}
