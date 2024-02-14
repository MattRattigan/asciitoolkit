package letters

import "testing"

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
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCII(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
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
		{"isEmpty", "", true},
		{"isSpecial", "!", false},
		{"isExtended", "Æ", false},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := IsASCIIAlphanumeric(tc.input)
			if result != tc.expected {
				t.Errorf("Input %s, Want %v, got %#v", tc.input, tc.expected, result)
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
		//{"isEmpty", 0, true},
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
}

func TestIsASCIILetter(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', true},
		{"isLower", 'a', true},
		//{"isEmpty", 0, true},
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
}

func TestIsASCIIUpper(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', true},
		{"isLower", 'a', false},
		//{"isEmpty", 0, true},
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
}

func TestIsASCIIPunct(t *testing.T) {
	t.Parallel()
	test := []testIs{
		{"isDigit", '1', false},
		{"isUpper", 'A', false},
		{"isLower", 'a', false},
		//{"isEmpty", 0, false},
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
}
