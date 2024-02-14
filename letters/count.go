package letters

// CountASCIIDigits counts the number of ASCII digit characters ('0'-'9') in the
// given string, return an int as the amount of counted digits.
func CountASCIIDigits(s string) int {
	count := 0
	for _, r := range s {
		if IsASCIIDigit(r) {
			count++
		}
	}
	return count
}

// CountASCIILetters counts the total number of ASCII letter characters (both
// uppercase and lowercase 'a' through 'z' and 'A' through 'Z') in the given
// string, return an int as the amount of counted letters.
func CountASCIILetters(s string) int {
	count := 0
	for _, r := range s {
		if IsASCIILetter(r) {
			count++
		}
	}
	return count
}

// CountASCIIUpper counts the number of uppercase ASCII letter characters ('A'
// through 'Z') return an int as the amount of counted uppercase letters, return
// an int as the amount of counted letters.
func CountASCIIUpper(s string) int {
	count := 0
	for _, r := range s {
		if IsASCIIUpper(r) {
			count++
		}
	}
	return count
}

// CountASCIILower counts the number of lowercase ASCII letter characters ('a'
// through 'z') in the given string, return an int as the amount of counted
// letters.
func CountASCIILower(s string) int {
	count := 0
	for _, r := range s {
		if IsASCIILower(r) {
			count++
		}
	}
	return count
}

// CountASCIIPunct counts the number of ASCII punctuation characters in the given
// string, return an int as the amount of counted characters.
func CountASCIIPunct(s string) int {
	count := 0
	for _, r := range s {
		if IsASCIIPunct(r) {
			count++
		}
	}
	return count
}
