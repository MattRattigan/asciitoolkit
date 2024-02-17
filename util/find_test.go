package util

import (
	"reflect"
	"testing"
)

type test struct {
	name  string
	input string
	want  []string
}

func TestFindASCIIDigits(t *testing.T) {
	tests := []test{
		{
			name:  "Only digits",
			input: "12345",
			want:  []string{"12345"},
		},

		{
			name:  "Mixed characters",
			input: "abc123def456",
			want:  []string{"123", "456"},
		},
		{
			name:  "Continuous and separated digits",
			input: "123abc456def 7890",
			want:  []string{"123", "456", "7890"},
		},

		{
			name:  "Digits with spaces",
			input: "123 456 789",
			want:  []string{"123", "456", "789"},
		},
	}

	for _, tc := range tests {
		got := FindASCIIDigits(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}

// TestFindASCIIWords tests the FindASCIIWords function for accuracy.
func TestFindASCIIWords(t *testing.T) {
	tests := []test{
		{
			name:  "Only words",
			input: "hello world",
			want:  []string{"hello", "world"},
		},
		{
			name:  "Mixed characters",
			input: "123abc def456",
			want:  []string{"abc", "def"},
		},
		{
			name:  "Words with numbers and punctuation",
			input: "hello, world! 1234 test123 test",
			want:  []string{"hello", "world", "test", "test"},
		},
		{
			name:  "All letters",
			input: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			want:  []string{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		},
	}

	for _, tc := range tests {
		got := FindASCIIWords(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
		}
	}
}
