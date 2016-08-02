package randchars

import (
	"testing"
)

func TestAlphaNum(t *testing.T) {
	g := NewGenerator()
	g.rng.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "AM"},
		{4, "p00A"},
		{10, "7cpFLjUAJM"},
	}
	for _, test := range tests {
		b := g.AlphaNum(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}

func TestAlpha(t *testing.T) {
	g := NewGenerator()
	g.rng.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "WY"},
		{4, "tuiY"},
		{10, "rcbnjhOmJy"},
	}
	for _, test := range tests {
		b := g.Alpha(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}
