package util

import "testing"

func TestFilterASCIIDigits(t *testing.T) {
	t.Parallel()
	input := "Hello, World! 12345"
	expected := "12345"
	if result := FilterASCIIDigits(input); result != expected {
		t.Errorf("FilterASCIIDigits(%q) = Got: %q, Want: %q", input, result, expected)
	}
}

func TestFilterASCIILetters(t *testing.T) {
	t.Parallel()
	input := "123 abc DEF!"
	expected := "abcDEF"
	if result := FilterASCIILetters(input); result != expected {
		t.Errorf("FilterASCIILetters(%q) = Got: %q, Want: %q", input, result, expected)
	}
}

func TestFilterASCIIUpper(t *testing.T) {
	t.Parallel()
	input := "Hello, World! 123"
	expected := "HW"
	if result := FilterASCIIUpper(input); result != expected {
		t.Errorf("FilterASCIIUpper(%q) = Got: %q, Want: %q", input, result, expected)
	}
}

func TestFilterASCIILower(t *testing.T) {
	t.Parallel()
	input := "Hello, World! 123"
	expected := "elloorld"
	if result := FilterASCIILower(input); result != expected {
		t.Errorf("FilterASCIILower(%q) = Got: %q, Want %q", input, result, expected)
	}
}

func TestFilterASCIIPunct(t *testing.T) {
	t.Parallel()
	input := "Hello, World! 123"
	expected := ",!"
	if result := FilterASCIIPunct(input); result != expected {
		t.Errorf("FilterASCIIPunct(%q) = Got: %q, Want: %q", input, result, expected)
	}
}
