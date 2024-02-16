package util

import (
	"fmt"
	"github.com/MattRattigan/asciitoolkit/ascii"
	"regexp"
)

func ToASCIIHex(s string) string {
	return ""
}

func FromASCIIHex(hexStr string) (string, error) {
	return "", nil
}

func MapASCIIFunc(s string, f func(rune) rune) string {
	return ""
}

func CompareASCIIIgnoreCase() {

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

func EncodeASCIIBase64() {

}
func DecodeASCIIBase64() {

}

func TrimASCIISpace() {

}

func ReplaceASCIIChars() {

}

// IsValidEmail provides basic email validation
func IsValidEmail(email string) bool {
	pattern := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}
func ASCIICharFrequency() {

}

func IsValidIdentifier(id string, min, max int) bool {
	pattern := fmt.Sprintf(`^\w{%d,%d}$`, min, max)
	re := regexp.MustCompile(pattern)
	return re.MatchString(id)
}

func toLower[T string | rune](v T) {
	switch val := any(v).(type) {
	case rune:
		// Find A
		_ = val
	}
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
