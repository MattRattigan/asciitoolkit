package ascii

import (
	"testing"
)

type testIs struct {
	name     string
	input    rune
	expected bool
}

func TestIsASCII(t *testing.T) {
	t.Parallel()
	test := []struct {
		name, input string
		expected    bool
	}{
		{"isDigit", "1", true},
		{"isUpper", "A", true},
		{"isLower", "a", true},
		{"isEmpty", "", true},
		{"isSpecial", "!", true},
		{"isExtended", "Æ", false},
		{"isExtended2", "¼", false},
		{"Whitespace", "  ", true},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCII(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}

	emptyAndStringTest := []struct {
		name, input string
		expected    bool
	}{
		{"isEmpty", "", true},
		{"String", "boomastic fantastic", true},
	}

	for _, tc := range emptyAndStringTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCII(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting empty string or single string character: got %v", tc.input)
			}
		})
	}
}

func TestIsASCIIAlphanumeric(t *testing.T) {
	t.Parallel()
	test := []struct {
		name, input string
		expected    bool
	}{
		{"isDigit", "1", true},
		{"isUpper", "A", true},
		{"isLower", "a", true},
		{"isEmpty", "", false},
		{"isSpecial", "!", false},
		{"isExtended", "Æ", false},
		{"String", "boomastic fantastic", false},
		{"Whitespace", "  ", false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIAlphanumeric(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
	// Test Empty and string with whitespace
	runeTest := []struct {
		name     string
		input    rune
		expected bool
	}{
		{"isDigit", '1', true},
		{"isUpper", 'A', true},
		{"isLower", 'a', true},
		{"isSpecial", '!', false},
		{"isExtended", 'Æ', false},
	}

	for _, tc := range runeTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIAlphanumeric(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting empty string or single string character: got %v", tc.input)
			}
		})
	}
}

func TestIsASCIIDigit(t *testing.T) {
	t.Parallel()

	test := []testIs{
		{"isDigit", '1', true},
		{"isUpper", 'A', false},
		{"isLower", 'a', false},
		{"isSpecial", '!', false},
		{"isExtended", 'Æ', false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIDigit(tc.input)
			if result != tc.expected {
				t.Errorf("Input %c, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}

	stringTest := []struct {
		name, input string
		expected    bool
	}{
		{"isEmpty", "", false},
		{"String", "larry", false},
		{"isDigit", "1", true},
		{"isUpper", "A", false},
		{"isLower", "a", false},
		{"isSpecial", "!", false},
		{"isExtended", "Æ", false},
		{"Whitespace", "  ", false},
	}

	for _, tc := range stringTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIDigit(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting string or rune: got empty string")
			}
		})
	}
}

func TestIsASCIILetter(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', true},
		{"isLower", 'a', true},
		{"isSpecial", '!', false},
		{"isExtended", 'Æ', false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIILetter(tc.input)
			if result != tc.expected {
				t.Errorf("Input %c, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}

	stringTest := []struct {
		name, input string
		expected    bool
	}{
		{"isEmpty", "", false},
		{"String", "larry", false},
		{"isDigit", "1", false},
		{"isUpper", "A", true},
		{"isLower", "a", true},
		{"isSpecial", "!", false},
		{"isExtended", "Æ", false},
		{"Whitespace", "  ", false},
	}

	for _, tc := range stringTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIILetter(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting string or rune: got empty string")
			}
		})
	}
}

func TestIsASCIIUpper(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', true},
		{"isLower", 'a', false},
		{"isSpecial", '!', false},
		{"isExtended", 'Æ', false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIUpper(tc.input)
			if result != tc.expected {
				t.Errorf("Input %c, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}

	stringTest := []struct {
		name, input string
		expected    bool
	}{
		{"isEmpty", "", false},
		{"String", "larry", false},
		{"isDigit", "1", false},
		{"isUpper", "A", true},
		{"isLower", "a", false},
		{"isSpecial", "!", false},
		{"isExtended", "Æ", false},
		{"Whitespace", "  ", false},
	}

	for _, tc := range stringTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIUpper(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting string or rune: got empty string")
			}
		})
	}
}

func TestIsASCIILower(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', false},
		{"isLower", 'a', true},
		//{"isEmpty", 0, true},
		{"isSpecial", '!', false},
		{"isExtended", 'Æ', false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIILower(tc.input)
			if result != tc.expected {
				t.Errorf("Input %c, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
	stringTest := []struct {
		name, input string
		expected    bool
	}{
		{"isEmpty", "", false},
		{"String", "larry", false},
		{"isDigit", "1", false},
		{"isUpper", "A", false},
		{"isLower", "a", true},
		{"isSpecial", "!", false},
		{"isExtended", "Æ", false},
		{"Whitespace", "  ", false},
	}

	for _, tc := range stringTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIILower(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting string or rune: got empty string")
			}
		})
	}
}

func TestIsASCIIPunct(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', false},
		{"isLower", 'a', false},
		{"isSpecial", '!', true},
		{"isExtended", 'Æ', false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIPunct(tc.input)
			if result != tc.expected {
				t.Errorf("Input %c, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
	stringTest := []struct {
		name, input string
		expected    bool
	}{
		{"isEmpty", "", false},
		{"String", "larry", false},
		{"isDigit", "1", false},
		{"isUpper", "A", false},
		{"isLower", "a", false},
		{"isSpecial", "!", true},
		{"isExtended", "Æ", false},
		{"Whitespace", "  ", false},
	}

	for _, tc := range stringTest {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIPunct(tc.input)
			if result != tc.expected {
				t.Errorf("Error expecting string or rune: got empty string")
			}
		})
	}
}

func TestIsASCIIWhitespace(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"Whitespace", ' ', true},
		{"Tab", '\t', true},
		{"NewLine", '\n', true},
		{"Carriage Return", '\r', true},
		{"Form Feed", '\f', true},
		{"Vertical Tab", '\v', true},
		{"A Rune", 'x', false},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIWhitespace(tc.input)
			if result != tc.expected {
				t.Errorf("Want: %v, Got: %v", result, tc.expected)
			}
		})
	}

	testString := []struct {
		name, input string
		expected    bool
	}{
		{"Whitespace", " ", true},
		{"Tab", "\t", true},
		{"NewLine", "\n", true},
		{"Carriage Return", "\r", true},
		{"Form Feed", "\f", true},
		{"Vertical Tab", "\v", true},
		{"A String", "x", false},
	}

	for _, tc := range testString {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIWhitespace(tc.input)
			if result != tc.expected {
				t.Errorf("Want: %v, Got: %v", result, tc.expected)
			}
		})
	}
}
