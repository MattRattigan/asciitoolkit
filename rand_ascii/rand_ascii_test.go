package rand_ascii

import (
	"github.com/matt/aciitools/ascii"
	"testing"
)

func TestRandASCIIDigit(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		got := RandASCIIDigit()
		if !ascii.IsASCIIDigit(got) {
			t.Errorf("expecting a rune, got: %c", got)
		}
	}

}

func TestRandASCIILetter(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		got := RandASCIILetter()
		if !ascii.IsASCIILetter(got) {
			t.Errorf("expecting a rune, got: %c", got)
		}
	}
}

func TestRandASCIILower(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		got := RandASCIILower()
		if !ascii.IsASCIILower(got) {
			t.Errorf("expecting a rune, got: %c", got)
		}
	}
}

func TestRandASCIIUpper(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		got := RandASCIIUpper()
		if !ascii.IsASCIIUpper(got) {
			t.Errorf("expecting a rune, got: %c", got)
		}
	}
}

func TestRandASCIIPunct(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		got := RandASCIIPunct()
		if !ascii.IsASCIIPunct(got) {
			t.Errorf("expecting a rune, got: %c", got)
		}
	}
}

func TestRandASCIIString(t *testing.T) {
	test := []struct {
		name          string
		length, input int
	}{
		{"LowerPunctuation8", 8, LOW | PUNCT},
		{"UpperDigit5", 5, UPP | DIG},
		{"UpperLowerPunctDigit10", 10, UPP | LOW | PUNCT | DIG},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := RandASCIIString(tc.length, tc.input)
			if !ascii.IsASCII(result) {
				t.Errorf("expecting string of comprised ASCII characters, got: %s", result)
			}
		})
	}
}
