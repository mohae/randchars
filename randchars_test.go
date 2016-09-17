package randchars

import (
	"fmt"
	mrand "math/rand"
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

func TestBase64(t *testing.T) {
	g := NewBase64()
	g.Seed(0)
	tests := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{2, "J6"},
		{4, "v8RG"},
		{10, "qQc0t5T5AQ"},
	}
	for _, test := range tests {
		b := g.Bytes(test.n)
		if string(b) != test.expected {
			t.Errorf("got %q; want %q", string(b), test.expected)
		}
	}
}

func BenchmarkMathRand_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MathRand(8)
	}
}

func BenchmarkAlpha_8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(8)
	}
}

func BenchmarkAlphaNum_8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(8)
	}
}

func BenchmarkLowerAlpha_8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(8)
	}
}

func BenchmarkLowerAlphaNum_8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(8)
	}
}

func BenchmarkUpperAlpha_8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(8)
	}
}

func BenchmarkUpperAlphaNum_8(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(8)
	}
}

func BenchmarkBase64_8(b *testing.B) {
	g := NewBase64()
	for i := 0; i < b.N; i++ {
		g.Bytes(8)
	}
}

func BenchmarkMathRand_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MathRand(16)
	}
}

func BenchmarkAlpha_16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(16)
	}
}

func BenchmarkAlphaNum_16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(16)
	}
}

func BenchmarkLowerAlpha_16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(16)
	}
}

func BenchmarkLowerAlphaNum_16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(16)
	}
}

func BenchmarkUpperAlpha_16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(16)
	}
}

func BenchmarkUpperAlphaNum_16(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(16)
	}
}

func BenchmarkBase64_16(b *testing.B) {
	g := NewBase64()
	for i := 0; i < b.N; i++ {
		g.Bytes(16)
	}
}

func BenchmarkMathRand_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MathRand(32)
	}
}

func BenchmarkAlpha_32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(32)
	}
}

func BenchmarkAlphaNum_32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(32)
	}
}

func BenchmarkLowerAlpha_32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(32)
	}
}

func BenchmarkLowerAlphaNum_32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(32)
	}
}

func BenchmarkUpperAlpha_32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(32)
	}
}

func BenchmarkUpperAlphaNum_32(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(32)
	}
}

func BenchmarkBase64_32(b *testing.B) {
	g := NewBase64()
	for i := 0; i < b.N; i++ {
		g.Bytes(32)
	}
}

func BenchmarkMathRand_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MathRand(64)
	}
}

func BenchmarkAlpha_64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.Alpha(64)
	}
}

func BenchmarkAlphaNum_64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(64)
	}
}

func BenchmarkLowerAlpha_64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(64)
	}
}

func BenchmarkLowerAlphaNum_64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(64)
	}
}

func BenchmarkUpperAlpha_64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(64)
	}
}

func BenchmarkUpperAlphaNum_64(b *testing.B) {
	g := NewGenerator()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(64)
	}
}

func BenchmarkBase64_64(b *testing.B) {
	g := NewBase64()
	for i := 0; i < b.N; i++ {
		g.Bytes(64)
	}
}

func MathRand(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, alphaNum[mrand.Int31n(int32(len(alphaNum)))])
	}
	return id
}
