package letters

import "testing"

type testContains struct {
	name, input string
	expected    bool
}

func TestContainsASCIIDigit(t *testing.T) {
	t.Parallel()
	test := []testContains{
		{"ContainsDigitsLower", "abc123", true},
		{"OnlyDigits", "123456", true},
		{"NoDigits", "abcdef", false},
		{"EmptyString", "", false},
		{"SpecialCharacters", "!@#$%^", false},
		{"DigitsAndSpecialChars", "123!@#", true},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsASCIIDigit(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestContainsASCIIUpper(t *testing.T) {
	t.Parallel()

	test := []testContains{
		{"UppercaseLetters", "AZFDSE", true},
		{"LowercaseLetters", "fapwer", false},
		{"ContainsDigitsLower", "abc123", false},
		{"LowerUpper", "sffRTFD", true},
		{"EmptyString", "", false},
		{"SpecialCharacters", "!@#$%^", false},
		{"DigitsAndSpecialChars", "123!@#", false},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsASCIIUpper(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestContainsASCIILetter(t *testing.T) {
	t.Parallel()
	test := []testContains{
		{"UppercaseLetters", "AZFDSE", true},
		{"LowercaseLetters", "fapwer", true},
		{"ContainsDigitsLower", "abc123", true},
		{"LowerUpper", "sffRTFD", true},
		{"EmptyString", "", false},
		{"SpecialCharacters", "!@#$%^", false},
		{"DigitsAndSpecialChars", "123!@#", false},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsASCIILetter(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestContainsASCIILower(t *testing.T) {
	t.Parallel()
	test := []testContains{
		{"UppercaseLetters", "AZFDSE", false},
		{"LowercaseLetters", "fapwer", true},
		{"ContainsDigitsLower", "abc123", true},
		{"ContainDigitsUpper", "ACD453", false},
		{"LowerUpper", "sffRTFD", true},
		{"EmptyString", "", false},
		{"SpecialCharacters", "!@#$%^", false},
		{"DigitsAndSpecialChars", "123!@#", false},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsASCIILower(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestContainsASCIIPunct(t *testing.T) {
	t.Parallel()
	test := []testContains{
		{"UppercaseLetters", "AZFDSE", false},
		{"LowercaseLetters", "fapwer", false},
		{"ContainsDigits", "abc123", false},
		{"LowerUpper", "sffRTFD", false},
		{"EmptyString", "", false},
		{"SpecialCharacters", "!@#$%^", true},
		{"DigitsAndSpecialChars", "123!@#", true},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ContainsASCIIPunct(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}
