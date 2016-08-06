package randchars

import (
	"testing"
)

func TestAlphaNum(t *testing.T) {
	g := NewGenerator()
	g.Seed(0)
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
	g.Seed(0)
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
	g.Seed(0)
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
	g.Seed(0)
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
	g.Seed(0)
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
	g.Seed(0)
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

func BenchmarkAlpha8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(8)
	}
}

func BenchmarkAlphaNum8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(8)
	}
}

func BenchmarkLowerAlpha8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(8)
	}
}

func BenchmarkLowerAlphaNum8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(8)
	}
}

func BenchmarkUpperAlpha8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(8)
	}
}

func BenchmarkUpperAlphaNum8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(8)
	}
}

func BenchmarkAlpha16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(16)
	}
}

func BenchmarkAlphaNum16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(16)
	}
}

func BenchmarkLowerAlpha16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(16)
	}
}

func BenchmarkLowerAlphaNum16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(16)
	}
}

func BenchmarkUpperAlpha16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(16)
	}
}

func BenchmarkUpperAlphaNum16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(16)
	}
}

func BenchmarkAlpha32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(32)
	}
}

func BenchmarkAlphaNum32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(32)
	}
}

func BenchmarkLowerAlpha32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(32)
	}
}

func BenchmarkLowerAlphaNum32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(32)
	}
}

func BenchmarkUpperAlpha32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(32)
	}
}

func BenchmarkUpperAlphaNum32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(32)
	}
}

func BenchmarkAlpha64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(64)
	}
}

func BenchmarkAlphaNum64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(64)
	}
}

func BenchmarkLowerAlpha64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(64)
	}
}

func BenchmarkLowerAlphaNum64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(64)
	}
}

func BenchmarkUpperAlpha64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(64)
	}
}

func BenchmarkUpperAlphaNum64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(64)
	}
}
