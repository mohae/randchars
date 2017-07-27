// Package randchars quickly generates a chunk of random ASCII characters
// using a PRNG. Two different generators are provide: Generator and Base64.
//
// Generator provides more flexibility in the set of characters used:
// a-zA-Z0-9, a-z0-9, A-Z0-9, a-zA-Z, a-z, A-Z, Base64, as defined in Table 1
// of RFC 4648, and Base64URL, as defined in Table 2 of RFC 4648.
//
// Base64 generates a chunk of Base 64 random characters. The character set
// used is from Table 1 of RFC 4648.
//
// Calls to the package functions using the package global genarator are
// threadsafe.
package randchars

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"

	pcg "github.com/dgryski/go-pcgr"
	xoro "github.com/dgryski/go-xoroshiro"
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
)

// use own
var gen *Generator
var mu sync.Mutex

// for base64 stuff
var genBase64 *Base64Generator
var mu64 sync.Mutex

func init() {
	gen = NewGenerator()
	genBase64 = NewBase64Generator()
}

// Generatorer is an interface for generators.
type Generatorer interface {
	AlphaNum(n int) []byte
	Alpha(n int) []byte
	LowerAlphaNum(n int) []byte
	LowerAlpha(n int) []byte
	UpperAlphaNum(n int) []byte
	UpperAlpha(n int) []byte
	Base64(n int) []byte
	Base64URL(n int) []byte
}

// Generator generates the random ASCII characters.  It relies on a PRNG that
// implements PCG: www.pcg-random.org.
type Generator struct {
	rng pcg.Rand
}

// Returns a seeded Generator that's ready to use.
func NewGenerator() *Generator {
	return &Generator{pcg.New(Int64(), 0)}
}

// NewGeneratorWithSeed a Generator using the received value as its seed.
func NewGeneratorWithSeed(seed int64) *Generator {
	return &Generator{pcg.New(seed, 0)}
}

// NewGeneratorSeedWithState a Generator using the received values as its seed
// and state.
func NewGeneratorSeedWithState(seed, state int64) *Generator {
	return &Generator{pcg.New(seed, state)}
}

// Seed seeds the Generator's prng.
func (g *Generator) Seed(n int64) {
	g.rng.Seed(n)
}

// SeedWithState seeds the Generator's prng and set's its state.
func (g *Generator) SeedWithState(seed, state int64) {
	g.rng.SeedWithState(seed, state)
}

// ReSeed seeds the Generator's prng using a value obtained from a CSPRNG.
func (g *Generator) ReSeed() {
	g.rng.Seed(Int64())
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

// Alpha returns a randomly generated []byte of length n using a-zA-Z. This
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

// LowerAlpha returns a randomly generated []byte of length n using a-z. This
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

// UpperAlpha returns a randomly generated []byte of length n using A-Z. This
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

// Base64 returns a randomly generated []byte of length n using Base64, as
// defined in Table 1 of RFC 4648. This will panic if n < 0.
func (g *Generator) Base64(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = base64[g.rng.Bound(uint32(len(base64)))]
	}
	return b
}

// Base64URL returns a randomly generated []byte of length n using Base64URL,
// as defined in Table 2 of RFC 4648. This will panic if n < 0.
func (g *Generator) Base64URL(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = base64URL[g.rng.Bound(uint32(len(base64URL)))]
	}
	return b
}

// Seed seeds the Generator's prng.
func Seed(n int64) {
	mu.Lock()
	gen.Seed(n)
	mu.Unlock()
}

// SeedWithState seeds the Generator's prng and set's its state.
func SeedWithState(seed, state int64) {
	mu.Lock()
	gen.SeedWithState(seed, state)
	mu.Unlock()
}

// ReSeed seeds the Generator's prng using a value obtained from a CSPRNG.
func ReSeed() {
	mu.Lock()
	gen.ReSeed()
	mu.Unlock()
}

// AlphaNum returns a randomly generated []byte of length n using a-zA-Z0-9.
// This will panic if n < 0.
func AlphaNum(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.AlphaNum(n)
}

// Alpha returns a randomly generated []byte of length n using a-zA-Z. This
// will panic if n < 0.
func Alpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.Alpha(n)
}

// LowerAlphaNum returns a randomly generated []byte of length n using a-z0-9.
// This will panic if n < 0.
func LowerAlphaNum(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.LowerAlphaNum(n)
}

// LowerAlpha returns a randomly generated []byte of length n using a-z. This
// will panic if n < 0.
func LowerAlpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.LowerAlpha(n)
}

// UpperAlphaNum returns a randomly generated []byte of length n using A-Z0-9.
// This will panic if n < 0.
func UpperAlphaNum(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.UpperAlphaNum(n)
}

// UpperAlpha returns a randomly generated []byte of length n using A-Z. This
// will panic if n < 0.
func UpperAlpha(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.UpperAlpha(n)
}

// Base64 returns a randomly generated []byte of length n using base64, as
// defined in Table 1 of RFC 4648. This will panic if n < 0.
func Base64(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.Base64(n)
}

// Base64URL returns a randomly generated []byte of length n using base64url,
// as defined in Table 2 of RFC 4648 filename safe base64. This will panic if
// n < 0.
func Base64URL(n int) []byte {
	mu.Lock()
	defer mu.Unlock()
	return gen.Base64URL(n)
}

// Base64 supports the Base 64 Alphabet as shown in Table 1 of RFC 4248. This
// uses an implementation of the XORoShiRo128+ PRNG:
// http://xoroshiro.di.unimi.it/.
//
// This is more performant than the other Generator that this package offers.
type Base64Generator struct {
	rng xoro.State
}

// NewBase64 returns an initialized Base64Generator that is ready to use. The
// seed value used is an int64 obtrained from a CSPRNG.
func NewBase64Generator() *Base64Generator {
	return &Base64Generator{xoro.New(Int64())}
}

// NewBase64GeneratorWithSeed a Base64Generator using the received value as
// its seed.
func NewBase64GeneratorWithSeed(seed int64) *Base64Generator {
	return &Base64Generator{xoro.New(seed)}
}

// Seed seeds Base64Generator's prng using the provided value.
func (g *Base64Generator) Seed(n int64) {
	g.rng.Seed(n)
}

// Reseed seeds Base64Generator's prng using a value obtained from a CSPRNG.
func (g *Base64Generator) Reseed() {
	g.rng.Seed(Int64())
}

// Bytes returns n randomly generated Base 64 bytes.
func (g *Base64Generator) Bytes(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, base64[g.rng.Int63n(int64(64))])
	}
	return id
}

// SeedBase64 seeds Base64Generator's prng using the provided value.
func SeedBase64(n int64) {
	mu64.Lock()
	genBase64.Seed(n)
	mu64.Unlock()
}

// ReseedBase64 seeds Base64Generator's prng using a value obtained from a CSPRNG.
func ReseedBase64() {
	mu64.Lock()
	genBase64.Reseed()
	mu64.Unlock()
}

// Base64Bytes returns n randomly generated Base 64 bytes.
func Base64Bytes(n int) []byte {
	mu64.Lock()
	defer mu64.Unlock()
	return genBase64.Bytes(n)
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
