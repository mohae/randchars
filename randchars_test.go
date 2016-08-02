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

func TestLowerAlphaNum(t *testing.T) {
	g := NewGenerator()
	g.rng.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "28"},
		{4, "tymk"},
		{10, "zkvf3tm2p2"},
	}
	for _, test := range tests {
		b := g.LowerAlphaNum(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}

func TestLowerAlpha(t *testing.T) {
	g := NewGenerator()
	g.rng.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "wy"},
		{4, "tuiy"},
		{10, "rcbnjhomjy"},
	}
	for _, test := range tests {
		b := g.LowerAlpha(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}

func TestUpperAlphaNum(t *testing.T) {
	g := NewGenerator()
	g.rng.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "28"},
		{4, "TYMK"},
		{10, "ZKVF3TM2P2"},
	}
	for _, test := range tests {
		b := g.UpperAlphaNum(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}

func TestUpperAlpha(t *testing.T) {
	g := NewGenerator()
	g.rng.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "WY"},
		{4, "TUIY"},
		{10, "RCBNJHOMJY"},
	}
	for _, test := range tests {
		b := g.UpperAlpha(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}
