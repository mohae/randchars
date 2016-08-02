// Package randchars generates random ASCII characters.  The set of characters
// can be one of the following: a-zA-Z0-9, a-z0-9, A-Z0-9, a-zA-Z, a-z, A-Z.
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
var gen Generator
var mu sync.Mutex

func init() {
	gen.rng.Seed(seed())
}

type Generator struct {
	rng pcg.Rand
}

// Returns a seeded Generator.
func NewGenerator() *Generator {
	var g Generator
	g.ReSeed()
	return &g
}

// ReSeed reseeds the Generator's prng.
func (g *Generator) ReSeed() {
	g.rng.Seed(seed())
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

func UpperAlpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.UpperAlpha(n)
}

// seed gets a random int64 using a CSPRNG.
func seed() int64 {
	bi := big.NewInt(1<<63 - 1)
	r, err := rand.Int(rand.Reader, bi)
	if err != nil {
		panic(fmt.Sprintf("entropy read error: %s", err))
	}
	return r.Int64()
}
