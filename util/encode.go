package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	//"github.com/MattRattigan/asciitoolkit/ascii"
	"regexp"
	"strings"
)

// ToASCIIHex converts a string to its ASCII hexadecimal representation.
func ToASCIIHex(s string) string {
	result := ""
	for _, c := range s {
		result += fmt.Sprintf("%02x", c)
	}
	return result
}

// FromASCIIHex converts an ASCII hexadecimal representation back to the original string.
func FromASCIIHex(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err // Return the error if the input is not a valid hexadecimal string.
	}
	return string(bytes), nil
}

// MapASCIIFunc maps a function to each rune in the string and returns a new string.
func MapASCIIFunc(s string, f func(rune) rune) string {
	result := strings.Builder{}
	for _, r := range s {
		result.WriteRune(f(r))
	}
	return result.String()
}

// CompareASCIIIgnoreCase compares two ASCII strings, ignoring case differences.
// Returns true if the strings are equal ignoring case, false otherwise.
func CompareASCIIIgnoreCase(s1, s2 string) bool {
	// Convert both strings to the same case for a case-insensitive comparison
	return strings.EqualFold(s1, s2)
}

// ToASCIILowerInPlace converts the input to lowercase in place. It supports both string and rune types.
func ToASCIILowerInPlace[T string | rune](v *T) {
	switch val := any(v).(type) {
	case *string:
		*val = ToLowerASCII(*val)
	case *rune:
		if ascii.IsASCIIUpper(*val) {
			*val = *val + ('a' - 'A')
		}
	}
}

// ToASCIIUpperInPlace converts the input to uppercase in place. It supports both string and rune types.
func ToASCIIUpperInPlace[T string | rune](v *T) {
	switch val := any(v).(type) {
	case *string:
		*val = ToUpperASCII(*val)
	case *rune:
		if ascii.IsASCIILower(*val) {
			*val = *val - ('a' - 'A')
		}
	}
}

// EncodeASCIIBase64 encodes a given ASCII string into a Base64 encoded string.
func EncodeASCIIBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// DecodeASCIIBase64 decodes a given Base64 encoded string back to its original ASCII form.
func DecodeASCIIBase64(encodedStr string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return "", err // Return the error if decoding fails
	}
	return string(bytes), nil
}

// TrimASCIISpace trims leading and trailing ASCII whitespace from a string.
func TrimASCIISpace(s string) string {
	return strings.TrimSpace(s)
}

// ReplaceASCIIChars replaces all occurrences of a specified character (oldChar) with another (newChar) in a string.
func ReplaceASCIIChars(s string, oldChar, newChar rune) string {
	return strings.ReplaceAll(s, string(oldChar), string(newChar))
}

// IsValidEmail provides basic email validation
func IsValidEmail(email string) bool {
	pattern := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// ASCIICharFrequency calculates the frequency of each ASCII character in the given string.
// It returns a slice where the index corresponds to the ASCII value of the character.
func ASCIICharFrequency(s string) [128]int {
	var frequency [128]int
	for _, char := range s {
		if char < 128 {
			frequency[char]++
		}
	}
	return frequency
}

func IsValidIdentifier(id string, min, max int) bool {
	pattern := fmt.Sprintf(`^\w{%d,%d}$`, min, max)
	re := regexp.MustCompile(pattern)
	return re.MatchString(id)
}

func ToLowerASCII[T string | rune](v T) string {
	switch val := any(v).(type) {
	case string:
		result := []byte(val)
		for i, b := range result {
			if ascii.IsASCIIUpper(rune(b)) {
				result[i] = b + ('a' - 'A')
			}
		}
		return string(result)
	case rune:
		result := byte(val)
		if ascii.IsASCIIUpper(val) {
			return string(result + ('a' - 'A'))
		}
		return string(result)
	default:
		return ""
	}
}

func ToUpperASCII[T string | rune](v T) string {
	switch val := any(v).(type) {
	case string:
		result := []byte(val)
		for i, b := range result {
			if ascii.IsASCIILower(rune(b)) {
				result[i] = b - ('a' - 'A')
			}
		}
		return string(result)
	case rune:
		result := byte(val)
		if ascii.IsASCIILower(val) {
			return string(result - ('a' - 'A'))
		}
		return string(result)
	default:
		return ""
	}

}
