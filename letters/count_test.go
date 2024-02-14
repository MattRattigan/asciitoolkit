package letters

import "testing"

type testCount struct {
	name, input string
	expected    int
}

func TestCountASCIIDigits(t *testing.T) {
	t.Parallel()
	test := []testCount{
		{"Letters", "AdgFgi", 0},
		{"Count6", "987456", 6},
		{"EmptyString", "", 0},
		{"DigitsLower", "abc123", 3},
		{"DigitsAndSpecialChars", "123!@#", 3},
		{"SpecialCharacters", "!@#$%^", 0},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := CountASCIIDigits(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestCountASCIILetters(t *testing.T) {
	t.Parallel()
	test := []testCount{
		{"Letters6", "AdgFgi", 6},
		{"DigitsOnly", "987456", 0},
		{"EmptyString", "", 0},
		{"DigitsLower", "abc123", 3},
		{"DigitsAndSpecialChars", "123!@#", 0},
		{"SpecialCharacters", "!@#$%^", 0},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := CountASCIILetters(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestCountASCIIUpper(t *testing.T) {
	t.Parallel()
	test := []testCount{
		{"Upper2", "AdgFgi", 2},
		{"DigitsOnly", "987456", 0},
		{"EmptyString", "", 0},
		{"DigitsUpper", "ABC123", 3},
		{"DigitsAndSpecialChars", "123!@#", 0},
		{"SpecialCharacters", "!@#$%^", 0},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := CountASCIIUpper(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestCountASCIILower(t *testing.T) {
	t.Parallel()
	test := []testCount{
		{"Lower4", "AdgFgi", 4},
		{"DigitsOnly", "987456", 0},
		{"EmptyString", "", 0},
		{"DigitsLower", "abc123", 3},
		{"DigitsAndSpecialChars", "123!@#", 0},
		{"SpecialCharacters", "!@#$%^", 0},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := CountASCIILower(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}

func TestCountASCIIPunct(t *testing.T) {
	t.Parallel()
	test := []testCount{
		{"Punct4", "@d%F^&", 4},
		{"DigitsOnly", "987456", 0},
		{"EmptyString", "", 0},
		{"DigitsLower", "abc123", 0},
		{"DigitsAndSpecialChars3", "123!@#", 3},
		{"SpecialCharacters6", "!@#$%^", 6},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := CountASCIIPunct(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
			}
		})
	}
}
