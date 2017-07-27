// Package crandchars generates a chunk of random ASCII characters using a
// CSPRNG. The supported ranges are: a-zA-Z0-9, a-z0-9, A-Z0-9, a-zA-Z, a-z,
// A-Z, and Base 64 as defined in Table 1 of RFC 4648.
//
// Calls to the package functions are threadsafe.
package crandchars

import (
	"crypto/rand"
	"fmt"
	"sync"
)

const (
	alphaNum      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alpha         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerAlphaNum = "abcdefghijklmnopqrstuvwxyz0123456789"
	lowerAlpha    = "abcdefghijklmnopqrstuvwxyz"
	upperAlphaNum = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	upperAlpha    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base64        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
	base64URL     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	CacheSize     = 4096 // The number of random bytes to cache.
)

var gen *Generator
var genMu sync.Mutex

func init() {
	gen = New()
}

// Generator handles generation of random chars.
type Generator struct {
	cache     []byte
	cacheSize int
	current   int
}

// New returns a Generator that uses the default CacheSize.
func New() *Generator {
	return NewGenerator(CacheSize)
}

// NewGenerator returns a generator with a cache of n random bytes.
func NewGenerator(n int) *Generator {
	g := Generator{cache: make([]byte, n), cacheSize: n}
	g.read()
	return &g
}

// AlphaNum returns a randomly generated []byte of length n using a-zA-Z0-9.
// This will panic if n < 0.
func (g *Generator) AlphaNum(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = alphaNum[g.intN(uint8(len(alphaNum)))]
	}
	return b
}

// Alpha returns a randomly generated []byte of length n using a-zA-Z. This
// will panic if n < 0.
func (g *Generator) Alpha(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = alpha[g.intN(uint8(len(alpha)))]
	}
	return b
}

// LowerAlphaNum returns a randomly generated []byte of length n using a-z0-9.
// This will panic if n < 0.
func (g *Generator) LowerAlphaNum(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = lowerAlphaNum[g.intN(uint8(len(lowerAlphaNum)))]
	}
	return b
}

// LowerAlpha returns a randomly generated []byte of length n using a-z. This
// will panic if n < 0
func (g *Generator) LowerAlpha(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = lowerAlpha[g.intN(uint8(len(lowerAlpha)))]
	}
	return b
}

// UpperAlphaNum returns a randomly generated []byte of length n using A-Z0-9.
// This will panic if n < 0.
func (g *Generator) UpperAlphaNum(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = upperAlphaNum[g.intN(uint8(len(upperAlphaNum)))]
	}
	return b
}

// UpperAlpha returns a randomly generated []byte of length n using A-Z. This
// will panic if n < 0.
func (g *Generator) UpperAlpha(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = upperAlpha[g.intN(uint8(len(upperAlpha)))]
	}
	return b
}

// Base64 returns a series of randomly generated Base64 bytes with the
// requested length. This will panic if n < 0.
func (g *Generator) Base64(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = base64[g.intN(uint8(len(base64)))]
	}
	return b
}

// Base64URL returns a series of randomly generated URL and Filename safe
// Base64 bytes with the requested length. This will panic if n < 0.
func (g *Generator) Base64URL(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = base64URL[g.intN(uint8(len(base64)))]
	}
	return b
}

// AlphaNum returns a randomly generated []byte of length n using a-zA-Z0-9.
// This will panic if n < 0.
func AlphaNum(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.AlphaNum(n)
}

// Alpha returns a randomly generated []byte of length n using a-zA-Z. This
// will panic if n < 0.
func Alpha(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.Alpha(n)
}

// LowerAlphaNum returns a randomly generated []byte of length n using a-z0-9.
// This will panic if n < 0.
func LowerAlphaNum(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.LowerAlphaNum(n)
}

// LowerAlpha returns a randomly generated []byte of length n using a-z. This
// will panic if n < 0.
func LowerAlpha(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.LowerAlpha(n)
}

// UpperAlphaNum returns a randomly generated []byte of length n using A-Z0-9.
// This will panic if n < 0.
func UpperAlphaNum(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.UpperAlphaNum(n)
}

// UpperAlpha returns a randomly generated []byte of length n using A-Z. This
// will panic if n < 0.
func UpperAlpha(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.UpperAlpha(n)
}

// Base64 returns a series of randomly generated Base 64 bytes with the
// requested length.  This will panic if n < 0.
func Base64(n int) []byte {
	genMu.Lock()
	defer genMu.Unlock()
	return gen.Base64(n)
}

// read fills the cache.
func (g *Generator) read() {
	_, err := rand.Read(g.cache)
	if err != nil {
		panic(err)
	}
}

// intN gets an unbiased value from the cache of random byte values.
func (g *Generator) intN(bound uint8) int {
	threshold := -bound % bound
	for {
		n := g.cache[g.current]
		g.current++
		// if we're at the end; replenish the cache
		if g.current >= g.cacheSize {
			g.read()
			g.current = 0
		}
		if n >= threshold {
			return int(n % bound)
		}
	}
}
