// Package randchars quickly generates a chunk of random ASCII characters.
// The set of characters can be one of the following: a-zA-Z0-9, a-z0-9,
// A-Z0-9, a-zA-Z, a-z, and A-Z.
//
// Calls to the package functions are threadsafe.
package randchars

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"

	pcg "github.com/dgryski/go-pcgr"
)

const alphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowerAlphaNum = "abcdefghijklmnopqrstuvwxyz0123456789"
const lowerAlpha = "abcdefghijklmnopqrstuvwxyz"
const upperAlphaNum = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const upperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// use own
var gen *Generator
var mu sync.Mutex

func init() {
	gen = NewGenerator()
}

// Generator generates the random ASCII characters.  It contains it's own
// PRNG.
type Generator struct {
	rng pcg.Rand
}

// Returns a seeded Generator that's ready to use.
func NewGenerator() *Generator {
	return &Generator{pcg.New(Int64(), 0)}
}

// Seed seeds the Generator's prng.
func (g *Generator) Seed(n int64) {
	g.rng.Seed(n)
}

// Seed seeds the Generator's prng.
func Seed(n int64) {
	mu.Lock()
	gen.Seed(n)
	mu.Unlock()
}

// SeedWithState seeds the Generator's prng and set's its state.
func (g *Generator) SeedWithState(seed, state int64) {
	g.rng.SeedWithState(seed, state)
}

// SeedWithState seeds the Generator's prng and set's its state.
func SeedWithState(seed, state int64) {
	mu.Lock()
	gen.SeedWithState(seed, state)
	mu.Unlock()
}

// ReSeed seeds the Generator's prng using a value obtained from a CSPRNG.
func (g *Generator) ReSeed() {
	g.rng.Seed(Int64())
}

// ReSeed seeds the Generator's prng using a value obtained from a CSPRNG.
func ReSeed() {
	mu.Lock()
	gen.ReSeed()
	mu.Unlock()
}

// AlphaNum returns a randomly generated []byte of length n using a-zA-Z0-9.
// This will panic if n < 0.
func (g *Generator) AlphaNum(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, alphaNum[g.rng.Bound(uint32(len(alphaNum)))])
	}
	return id
}

// AlphaNum returns a randomly generated []byte of length n using a-zA-Z0-9.
// This is a thread-safe call; will panic if n < 0.
func AlphaNum(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.AlphaNum(n)
}

// Alpha returns a randomly generated []byte of length n using a-zA-Z.  This
// will panic if n < 0.
func (g *Generator) Alpha(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, alpha[g.rng.Bound(uint32(len(alpha)))])
	}
	return id
}

// Alpha returns a randomly generated []byte of length n using a-zA-Z.  This
// is a thread-safe call; will panic if n < 0.
func Alpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.Alpha(n)
}

// LowerAlphaNum returns a randomly generated []byte of length n using a-z0-9.
// This will panic if n < 0.
func (g *Generator) LowerAlphaNum(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, lowerAlphaNum[g.rng.Bound(uint32(len(lowerAlphaNum)))])
	}
	return id
}

// LowerAlphaNum returns a randomly generated []byte of length n using a-z0-9.
// This is a thread-safe call; will panic if n < 0.
func LowerAlphaNum(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.LowerAlphaNum(n)
}

// LowerAlpha returns a randomly generated []byte of length n using a-z.  This
// will panic if n < 0.
func (g *Generator) LowerAlpha(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, lowerAlpha[g.rng.Bound(uint32(len(lowerAlpha)))])
	}
	return id
}

// LowerAlpha returns a randomly generated []byte of length n using a-z.  This
// is a thread-safe call; will panic if n < 0.
func LowerAlpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.LowerAlpha(n)
}

// UpperAlphaNum returns a randomly generated []byte of length n using A-Z0-9.
// This will panic if n < 0.
func (g *Generator) UpperAlphaNum(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, upperAlphaNum[g.rng.Bound(uint32(len(upperAlphaNum)))])
	}
	return id
}

// UpperAlphaNum returns a randomly generated []byte of length n using A-Z0-9.
// This is a thread-safe call; will panic if n < 0.
func UpperAlphaNum(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.UpperAlphaNum(n)
}

// UpperAlpha returns a randomly generated []byte of length n using A-Z.  This
// will panic if n < 0.
func (g *Generator) UpperAlpha(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, upperAlpha[g.rng.Bound(uint32(len(upperAlpha)))])
	}
	return id
}

// UpperAlpha returns a randomly generated []byte of length n using A-Z.  This
// is a thread-safe call; will panic if n < 0.
func UpperAlpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.UpperAlpha(n)
}

// Int64 gets an int64 value from a CSPRNG.
func Int64() int64 {
	bi := big.NewInt(1<<63 - 1)
	r, err := rand.Int(rand.Reader, bi)
	if err != nil {
		panic(fmt.Sprintf("entropy read error: %s", err))
	}
	return r.Int64()
}
