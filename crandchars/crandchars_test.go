package crandchars

import (
	"testing"
)

func BenchmarkAlphaNum_8(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(8)
	}
}

func BenchmarkLowerAlpha_8(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(8)
	}
}

func BenchmarkLowerAlphaNum_8(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(8)
	}
}

func BenchmarkUpperAlpha_8(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(8)
	}
}

func BenchmarkUpperAlphaNum_8(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(8)
	}
}

func BenchmarkBase64_8(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Base64(8)
	}
}

func BenchmarkAlpha_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Alpha(16)
	}
}

func BenchmarkAlphaNum_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(16)
	}
}

func BenchmarkLowerAlpha_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(16)
	}
}

func BenchmarkLowerAlphaNum_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(16)
	}
}

func BenchmarkUpperAlpha_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(16)
	}
}

func BenchmarkUpperAlphaNum_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(16)
	}
}

func BenchmarkBase64_16(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Base64(16)
	}
}

func BenchmarkAlpha_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Alpha(32)
	}
}

func BenchmarkAlphaNum_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(32)
	}
}

func BenchmarkLowerAlpha_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(32)
	}
}

func BenchmarkLowerAlphaNum_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(32)
	}
}

func BenchmarkUpperAlpha_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(32)
	}
}

func BenchmarkUpperAlphaNum_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(32)
	}
}

func BenchmarkBase64_32(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Base64(32)
	}
}

func BenchmarkAlpha_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Alpha(64)
	}
}

func BenchmarkAlphaNum_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.AlphaNum(64)
	}
}

func BenchmarkLowerAlpha_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlpha(64)
	}
}

func BenchmarkLowerAlphaNum_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.LowerAlphaNum(64)
	}
}

func BenchmarkUpperAlpha_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlpha(64)
	}
}

func BenchmarkUpperAlphaNum_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.UpperAlphaNum(64)
	}
}

func BenchmarkBase64_64(b *testing.B) {
	g := New()
	for i := 0; i < b.N; i++ {
		g.Base64(64)
	}
}
