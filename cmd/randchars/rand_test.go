package main

import (
	"testing"

	"github.com/mohae/randchars"
)

func TestRandGen(t *testing.T) {
	tests := []struct {
		chars    string
		n        int
		expected string
		err      string
	}{
		{"", 10, "", "\"\" is not supported"},
		{"alphanum", 0, "", "0: invalid character amount; must be > 0"},
		{"alphanum", 12, "AMp00A7cpFLj", ""},
		{"alpha", 12, "WYtuiYrcbnjh", ""},
		{"loweralphanum", 12, "28tymkzkvf3t", ""},
		{"loweralpha", 12, "wytuiyrcbnjh", ""},
		{"upperalphanum", 12, "28TYMKZKVF3T", ""},
		{"upperalpha", 12, "WYTUIYRCBNJH", ""},
		{"base64", 12, "iw7GmwTUjnz7", ""},
	}
	for _, test := range tests {
		g, err := NewGenerator(test.n, false, test.chars)
		if err != nil {
			if err.Error() != test.err {
				t.Errorf("%s: got %q; want %q", test.chars, err, test.err)
			}
			continue
		}
		g.Gen.(*randchars.Generator).Seed(0)
		b := g.GetChars(test.n)
		if string(b) != test.expected {
			t.Errorf("%s: got %q; want %q", test.chars, string(b), test.expected)
		}

	}
}
