package rand_ascii

import (
	"github.com/MattRattigan/asciitoolkit/ascii"
	"testing"
)

func TestRandDigit(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		for _, v := range []any{RandCharDigitStr(), RandASCIIDigitRune()} {
			switch got := v.(type) {
			case rune:
				if !ascii.IsASCIIDigit(got) {
					t.Errorf("expecting rune, got %c", got)
				}
			case string:
				if !ascii.IsASCIIDigit(got) {
					t.Errorf("expecting single character string, got: %s", got)
				}
			}
		}
	}
}

func TestRandLetter(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		for _, v := range []any{RandASCIILetterStr(), RandASCIILetterRune()} {
			switch got := v.(type) {
			case rune:
				if !ascii.IsASCIILetter(got) {
					t.Errorf("expecting rune, got %c", got)
				}
			case string:
				if !ascii.IsASCIILetter(got) {
					t.Errorf("expecting single character string, got: %s", got)
				}
			}
		}
	}
}

func TestRandUpper(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		for _, v := range []any{RandASCIIUpperStr(), RandASCIIUpperRune()} {
			switch got := v.(type) {
			case rune:
				if !ascii.IsASCIIUpper(got) {
					t.Errorf("expecting rune, got %c", got)
				}
			case string:
				if !ascii.IsASCIIUpper(got) {
					t.Errorf("expecting single character string, got: %s", got)
				}
			}
		}
	}
}

func TestRandLower(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		for _, v := range []any{RandASCIILowerStr(), RandASCIILowerRune()} {
			switch got := v.(type) {
			case rune:
				if !ascii.IsASCIILower(got) {
					t.Errorf("expecting rune, got %c", got)
				}
			case string:
				if !ascii.IsASCIILower(got) {
					t.Errorf("expecting single character string, got: %s", got)
				}
			}
		}
	}
}
func TestRandPunct(t *testing.T) {
	t.Parallel()
	for i := 0; i < 1000; i++ {
		for _, v := range []any{RandASCIIPunctStr(), RandASCIIPunctRune()} {
			switch got := v.(type) {
			case rune:
				if !ascii.IsASCIIPunct(got) {
					t.Errorf("expecting rune, got %c", got)
				}
			case string:
				if !ascii.IsASCIIPunct(got) {
					t.Errorf("expecting single character string, got: %s", got)
				}
			}
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
		{"PunctForCoverage", 5, PUNCT},
		{"UpperForCoverage", 11, UPP},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := RandASCIIString(tc.length, tc.input)
			if !ascii.IsASCII(result) {
				t.Errorf("expecting string of comprised ASCII characters, got: %v", result)
			}
		})
	}
}
