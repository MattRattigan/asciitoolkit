package util

import (
	"testing"
)

func TestToASCIILowerInPlace(t *testing.T) {
	t.Parallel()
	test := []struct {
		name, got, want string
	}{
		{"Mixed Case", "Larry", "larry"},
		{"Uppercase", "Smokey", "smokey"},
		{"Lowercase", "Hungry", "hungry"},
		{"Empty", "", ""},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			ToASCIILowerInPlace(&tc.got)
			if tc.got != tc.want {
				t.Errorf("Want: %s, Got: %s", tc.want, tc.got)
			}
		})
	}

	testRune := []struct {
		name      string
		got, want rune
	}{
		{"Mixed Case", 'L', 'l'},
		{"Uppercase", 'S', 's'},
		{"Lowercase", 'H', 'h'},
		{"WhiteSpace", ' ', ' '},
	}

	for _, tc := range testRune {
		t.Run(tc.name, func(t *testing.T) {
			ToASCIILowerInPlace(&tc.got)
			if tc.got != tc.want {
				t.Errorf("Want: %v, Got: %v", tc.want, tc.got)
			}
		})
	}
}

func TestToASCIIUpperInPlace(t *testing.T) {
	t.Parallel()
	test := []struct {
		name, got, want string
	}{
		{"Mixed Case", "Larry", "LARRY"},
		{"Uppercase", "Smokey", "SMOKEY"},
		{"Lowercase", "Hungry", "HUNGRY"},
		{"Empty", "", ""},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			ToASCIIUpperInPlace(&tc.got)
			if tc.got != tc.want {
				t.Errorf("Want: %s, Got: %s", tc.want, tc.got)
			}
		})
	}

	testRune := []struct {
		name      string
		got, want rune
	}{
		{"Uppercase", 'A', 'A'},
		{"Lowercase", 'a', 'A'},
		{"Whitespace", ' ', ' '},
	}

	for _, tc := range testRune {
		t.Run(tc.name, func(t *testing.T) {
			ToASCIIUpperInPlace(&tc.got)
			if tc.got != tc.want {
				t.Errorf("Want: %c, Got: %c", tc.want, tc.got)
			}
		})
	}
}

func TestToLowerCase(t *testing.T) {
	t.Parallel()
	test := []struct {
		name, input, want string
	}{
		{"Mixed Case", "MeRry HaD a LiTtle Lamb", "merry had a little lamb"},
		{"Uppercase", "BELGIAN WAFFLES", "belgian waffles"},
		{"Lowercase", "bacon", "bacon"},
		{"Empty", "", ""},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ToLowerASCII(tc.input)
			if result != tc.want {
				t.Errorf("Want: %s, Got: %s", tc.want, result)
			}
		})
	}

	testRune := []struct {
		name           string
		want, expected rune
	}{
		{"Uppercase", 'A', 'a'},
		{"Lowercase", 'a', 'a'},
		{"Whitespace", ' ', ' '},
	}

	for _, tc := range testRune {
		t.Run(tc.name, func(t *testing.T) {
			result := ToLowerASCII(tc.want)
			if result != string(tc.expected) {
				t.Errorf("Want: %c, Got: %s", tc.want, result)
			}
		})
	}
}

func TestToUpperCaseASCII(t *testing.T) {
	t.Parallel()
	test := []struct {
		name, input, want string
	}{
		{"Mixed Case", "thE qUIck brOwn fOx jUmps OvEr thE LAzy dOg", "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"},
		{"Lowercase", "beef patty", "BEEF PATTY"},
		{"Uppercase", "TING", "TING"},
		{"Empty", "", ""},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ToUpperASCII(tc.input)
			if result != tc.want {
				t.Errorf("Want: %v, Got: %v", tc.want, result)
			}
		})
	}

	testRune := []struct {
		name           string
		want, expected rune
	}{
		{"Uppercase", 'A', 'A'},
		{"Lowercase", 'a', 'A'},
		{"Whitespace", ' ', ' '},
	}

	for _, tc := range testRune {
		t.Run(tc.name, func(t *testing.T) {
			result := ToUpperASCII(tc.want)
			if result != string(tc.expected) {
				t.Errorf("Want: %v, Got: %v", tc.expected, result)
			}
		})
	}
}
