package ascii

import (
	"strings"
	"testing"
)

// TestASCIILowercase tests the ASCIILowercase function to ensure it returns the correct string of lowercase ASCII ascii.
func TestASCIILowercase(t *testing.T) {
	t.Parallel()
	want := "abcdefghijklmnopqrstuvwxyz"
	got := ASCIILowercase()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

// TestASCIIUppercase tests the ASCIIUppercase function to ensure it returns the correct string of uppercase ASCII ascii.
func TestASCIIUppercase(t *testing.T) {
	t.Parallel()
	want := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	got := ASCIIUppercase()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

// TestASCIIPunctuation tests the TestASCIIPunctuation function to ensure it returns the correct string of punctuation ASCII ascii.
func TestASCIIPunctuation(t *testing.T) {
	t.Parallel()
	want := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	got := ASCIIPunctuation()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestASCIIHexDigits(t *testing.T) {
	t.Parallel()
	want := "0123456789abcdefABCDEF"
	got := HexDigits()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestASCIIOctDigits(t *testing.T) {
	t.Parallel()
	want := "01234567"
	got := OctDigits()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestPrintable(t *testing.T) {
	t.Parallel()
	var builder strings.Builder
	for _, i := range []string{Digits(), ASCIIUppercase(), ASCIILowercase(), ASCIIPunctuation(), Whitespace()} {
		builder.WriteString(i)
	}
	want := builder.String()
	got := Printable()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestWhitespace(t *testing.T) {
	t.Parallel()
	want := " \t\n\r\f\v"
	got := Whitespace()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestASCIILetters(t *testing.T) {
	t.Parallel()
	var builder strings.Builder
	for _, i := range []string{ASCIILowercase(), ASCIIUppercase()} {
		builder.WriteString(i)
	}
	want := builder.String()
	got := ASCIILetters()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}

func TestDigits(t *testing.T) {
	t.Parallel()
	want := "0123456789"
	got := Digits()

	if got != want {
		t.Errorf("want: %s, got %s", want, got)
	}
}
