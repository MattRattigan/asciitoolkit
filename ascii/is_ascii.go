package ascii

func isHelper[T Textual, U comparable](v T, lowBound, upBound U) bool {
	switch val := any(v).(type) {
	case rune:
		if ub, ok := any(upBound).(rune); ok {
			if lb, ok := any(lowBound).(rune); ok {
				return val >= lb && val <= ub
			}
		}
	case string:
		if len(val) == 1 {
			if ub, ok := any(upBound).(byte); ok {
				if lb, ok := any(lowBound).(byte); ok {
					return val[0] >= lb && val[0] <= ub
				}
			}
		}
	}
	return false
}

// IsASCIIDigit checks if the given rune or single character string v represents an ASCII digit between ('0' through '9').
// returns true if v is within the range of ASCII digits, false if otherwise.
func IsASCIIDigit[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return isHelper(val, '0', '9')
	case string:
		return isHelper(val, byte('0'), byte('9'))
	}
	return false
}

// IsASCIILetter checks if the given rune or single character string v represents an ASCII that reside between ('A' through 'Z') or .
// ('a' through 'z') returns true if v is within the range of ASCII, false if otherwise.
func IsASCIILetter[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return isHelper(val, 'A', 'Z') || isHelper(val, 'a', 'z')
	case string:
		return isHelper(val, byte('A'), byte('Z')) || isHelper(val, byte('a'), byte('z'))
	}
	return false
}

// IsASCIIUpper checks if the given rune r represents an ASCII uppercase ascii ('A' through 'Z').
// returns true if r is within the range of ASCII uppercase ascii, false if otherwise.
func IsASCIIUpper[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return isHelper(val, 'A', 'Z')
	case string:
		return isHelper(val, byte('A'), byte('Z'))
	}
	return false
}

// IsASCIILower checks if the given rune r represents an ASCII lowercase letter ('a' through 'z').
// returns true if r is within the range of ASCII lowercase ascii, false if otherwise.
func IsASCIILower[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return isHelper(val, 'a', 'z')
	case string:
		return isHelper(val, byte('a'), byte('z'))
	}
	return false
}

// IsASCIIPunct checks if the given rune r represents an ASCII punctuation character.
// ASCII punctuation characters are in the ranges of 33-47, 58-64, 91-96, and 123-126.
// It returns true if r is within any of these ranges, indicating it is a punctuation character; false otherwise.
func IsASCIIPunct[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return (val >= 33 && val <= 47) || (val >= 58 && val <= 64) ||
			(val >= 91 && val <= 96) || (val >= 123 && val <= 126)
	case string:
		if len(val) == 1 {
			return (val[0] >= 33 && val[0] <= 47) || (val[0] >= 58 && val[0] <= 64) ||
				(val[0] >= 91 && val[0] <= 96) || (val[0] >= 123 && val[0] <= 126)
		}
		return false
	default:
		return false
	}
}

// IsASCII checks whether character within the provided string are the standard ASCII characters
// (Printable and Control characters), "" will return true.
func IsASCII[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		if val > 127 {
			return false
		}
	case string:
		for _, r := range val {
			if r > 127 {
				return false
			}
		}
	default:
		return false
	}
	return true
}

// IsASCIIAlphanumeric checks whether the provided string is a set of alphanumeric characters.
func IsASCIIAlphanumeric[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return IsASCIIDigit(val) || IsASCIILetter(val)
	case string:
		// Return false for an empty string
		if val == "" {
			return false
		}
		for _, r := range val {
			if !IsASCIIDigit(r) && !IsASCIILetter(r) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func IsASCIIWhitespace[T Text](v T) bool {
	switch val := any(v).(type) {
	case rune:
		return val == ' ' || val == '\t' || val == '\n' || val == '\r' || val == '\f' || val == '\v'
	case string:
		if len(val) == 1 {
			r := rune(val[0])
			return r == ' ' || r == '\t' || r == '\n' || r == '\r' || r == '\f' || r == '\v'
		}
		return false
	default:
		return false
	}
}
