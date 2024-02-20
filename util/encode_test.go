package util

import (
	"strings"
	"testing"
)

type TestMultiple struct {
	name  string
	input interface{}
	want  string
}

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

	testMultiple := []struct {
		name  string
		input interface{}
		want  interface{}
	}{
		{"Mixed Case", "thE qUIck brOwn fOx jUmps OvEr thE LAzy dOg", "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"},
		{"Lowercase", "beef patty", "BEEF PATTY"},
		{"Uppercase", "TING", "TING"},
		{"Empty", "", ""},
		{"Uppercase", 'A', 'A'},
		{"Lowercase", 'a', 'A'},
		{"Whitespace", ' ', ' '},
	}

	for _, tc := range testMultiple {
		t.Run(tc.name, func(t *testing.T) {
			switch input := tc.input.(type) {
			case string:
				upper := ToUpperASCII(input)
				_ = upper
			case rune:
				upper := ToUpperASCII(input)
				_ = upper
			}
		})
	}

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

func TestToASCIIHex(t *testing.T) {
	t.Parallel()
	tests := []TestMultiple{
		{"String", "To the moon!", "546f20746865206d6f6f6e21"},
		{"SliceRune", []rune{'T', 'o', ' ', 't', 'h', 'e', ' ', 'm', 'o', 'o', 'n', '!'}, "546f20746865206d6f6f6e21"},
		{"MultipleStrings", []string{"Hello", ", World!"}, "48656c6c6f2c20576f726c6421"},
		{"SingleRune", 'T', "54"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var hexStr string
			switch input := tc.input.(type) {
			case string:
				hexStr = ToASCIIHex(input)
			case []rune:
				hexStr = ToASCIIHex(input)
			case rune:
				hexStr = ToASCIIHex(input)
			case []string:
				hexStr = ToASCIIHex(input)
			default:
				t.Fatalf("Unsupported type in test case")
			}

			if hexStr != tc.want {
				t.Errorf("Expected %s, got %s", tc.want, hexStr)
			}
		})
	}
}

// TestFromASCIIHex tests the FromASCIIHex function with various inputs.
func TestFromASCIIHex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		hexStr  string
		want    string
		wantErr bool
	}{
		{name: "empty string", hexStr: "", want: "", wantErr: false},
		{name: "valid hex string", hexStr: "68656c6c6f", want: "hello", wantErr: false},
		{name: "valid hex string with uppercase", hexStr: "48656C6C6F", want: "Hello", wantErr: false},
		{name: "invalid hex string", hexStr: "zzzz", want: "", wantErr: true},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromASCIIHex(tt.hexStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromASCIIHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromASCIIHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestCompareASCIIIgnoreCase tests the CompareASCIIIgnoreCase function with various inputs.
func TestCompareASCIIIgnoreCase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{name: "equal strings, same case", s1: "hello", s2: "hello", want: true},
		{name: "equal strings, different cases", s1: "Hello", s2: "hello", want: true},
		{name: "different strings", s1: "Hello", s2: "World", want: false},
		{name: "one empty string", s1: "", s2: "hello", want: false},
		{name: "both empty strings", s1: "", s2: "", want: true},
		{name: "strings with spaces", s1: "hello world", s2: "Hello World", want: true},
		{name: "strings with numbers", s1: "123", s2: "123", want: true},
		{name: "strings with special characters", s1: "!@#", s2: "!@#", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareASCIIIgnoreCase(tt.s1, tt.s2); got != tt.want {
				t.Errorf("CompareASCIIIgnoreCase(%v, %v) = %v, want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestMapASCIIFunc(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		mapper func(rune) rune
		want   string
	}{
		{
			name: "To uppercase", input: "hello",
			mapper: func(r rune) rune {
				return rune(strings.ToUpper(string(r))[0])
			},
			want: "HELLO",
		},
		{
			name: "To lowercase", input: "HELLO",
			mapper: func(r rune) rune {
				return rune(strings.ToLower(string(r))[0])
			},
			want: "hello",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapASCIIFunc(tt.input, tt.mapper)
			if got != tt.want {
				t.Errorf("MapASCIIFunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeASCIIBase64(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Encode empty string",
			input: "",
			want:  "",
		},
		{
			name:  "Encode simple ASCII",
			input: "hello",
			want:  "aGVsbG8=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeASCIIBase64(tt.input)
			if got != tt.want {
				t.Errorf("EncodeASCIIBase64(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDecodeASCIIBase64(t *testing.T) {
	t.Parallel()
	cases := []struct {
		encoded string
		want    string
		wantErr bool
	}{
		{"aGVsbG8=", "hello", false},
		{"", "", false},
		{"IV@#", "", true}, // Invalid Base64
	}

	for _, c := range cases {
		got, err := DecodeASCIIBase64(c.encoded)
		if (err != nil) != c.wantErr {
			t.Errorf("DecodeASCIIBase64(%q) unexpected error state: %v", c.encoded, err)
		}
		if got != c.want {
			t.Errorf("DecodeASCIIBase64(%q) = %q, want %q", c.encoded, got, c.want)
		}
	}
}

func TestTrimASCIISpace(t *testing.T) {
	t.Parallel()
	cases := []struct {
		given string
		want  string
	}{
		{"  hello  ", "hello"},
		{"hello world", "hello world"},
		{"", ""},
	}

	for _, c := range cases {
		got := TrimASCIISpace(c.given)
		if got != c.want {
			t.Errorf("TrimASCIISpace(%q) = %q, want %q", c.given, got, c.want)
		}
	}
}

func TestReplaceASCIIChars(t *testing.T) {
	t.Parallel()
	cases := []struct {
		s    string
		old  rune
		new  rune
		want string
	}{
		{"hello", 'l', 'w', "hewwo"},
		{"hello", 'a', 'x', "hello"}, // 'a' not in "hello"
		{"", 'x', 'y', ""},           // Empty string
	}

	for _, c := range cases {
		got := ReplaceASCIIChars(c.s, c.old, c.new)
		if got != c.want {
			t.Errorf("ReplaceASCIIChars(%q, %q, %q) = %q, want %q", c.s, c.old, c.new, got, c.want)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	t.Parallel()
	cases := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true},
		{"invalid@ email.com", false},
		{"", false},
	}

	for _, c := range cases {
		got := IsValidEmail(c.email)
		if got != c.want {
			t.Errorf("IsValidEmail(%q) = %t, want %t", c.email, got, c.want)
		}
	}
}

func TestASCIICharFrequency(t *testing.T) {
	t.Parallel()
	cases := []struct {
		s    string
		want [128]int
	}{
		{"hello", func() [128]int { var f [128]int; f['h'], f['e'], f['l'], f['o'] = 1, 1, 2, 1; return f }()},
		{"", [128]int{}},
	}

	for _, c := range cases {
		got := ASCIICharFrequency(c.s)
		for i, freq := range got {
			if freq != c.want[i] {
				t.Errorf("ASCIICharFrequency(%q) did not match expected frequency at %d", c.s, i)
				break
			}
		}
	}
}

func TestIsValidIdentifier(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		id   string
		min  int
		max  int
		want bool
	}{
		{"Valid Short ID", "abc", 3, 5, true},
		{"Valid Max Length ID", "abcde", 3, 5, true},
		{"Too Short", "ab", 3, 5, false},
		{"Max Only", "abcd", 0, 4, true},
		{"Unicode", "日本語", 2, 4, false}, // Assuming non-ASCII characters are not considered valid
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidIdentifier(tt.id, tt.min, tt.max)
			if got != tt.want {
				t.Errorf("IsValidIdentifier(%q, %d, %d) = %v, want %v", tt.id, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
