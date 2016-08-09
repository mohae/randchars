// Package randchars quickly generates a chunk of random ASCII characters.
// Two different generators are provide: Generator and Gen64.
//
// Generator provides more flexibility in the set of characters used:
// a-zA-Z0-9, a-z0-9, A-Z0-9, a-zA-Z, a-z, and A-Z.
//
// Gen64 provides a 64 character set: a-z, A-Z, 0-9, =, and _ (underscore).
// The time to generate n random characters is ~20% less than Generator's.
//
// Calls to the package functions are threadsafe.
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
	ascii64       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789=_"
)

// use own
var gen *Generator
var mu sync.Mutex

// for ascii64 stuff
var gen64 *Gen64
var mu64 sync.Mutex

func init() {
	gen = NewGenerator()
	gen64 = NewGen64()
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

// Gen64 only supports an ASCII chracter set that's 64 bytes in length:
// a-z, A-Z, 0-9, =, and _ (underscore).  This is to take advantage of a
// shortcut that can be done on bounds that are powers of 2, which eliminates
// most of the cost of handling modulo bias.  Gen64 uses a PRNG that implements
// XORoShiRo128+: http://xoroshiro.di.unimi.it/.
type Gen64 struct {
	rng xoro.State
}

// NewGen64 returns an initialized Gen64 that is ready to use.  The seed value
// used is an int64 generated by a CSPRNG.
func NewGen64() *Gen64 {
	return &Gen64{xoro.New(Int64())}
}

// NewGener64WithSeed a Gen64 Generator using the received value as its seed.
func NewGen64WithSeed(seed int64) *Gen64 {
	return &Gen64{xoro.New(seed)}
}

// Seed seeds Gen64's prng.
func (g *Gen64) Seed(n int64) {
	g.rng.Seed(n)
}

// Seed seeds Gen64's prng.
func Seed64(n int64) {
	mu64.Lock()
	gen64.Seed(n)
	mu64.Unlock()
}

// ReSeed seeds Gen64's prng using a value obtained from a CSPRNG.
func (g *Gen64) ReSeed() {
	g.rng.Seed(Int64())
}

// ReSeed seeds Gen64's prng using a value obtained from a CSPRNG.
func ReSeed64() {
	mu64.Lock()
	gen64.ReSeed()
	mu64.Unlock()
}

// Bytes returns a series of randomly generated ASCII64 bytes with the
// requested length.
func (g *Gen64) Bytes(n int) []byte {
	if n < 0 {
		panic(fmt.Sprintf("%d: value out of bounds", n))
	}
	id := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		id = append(id, ascii64[g.rng.Int63n(int64(64))])
	}
	return id
}

// Bytes returns a series of randomly generated ASCII64 bytes with the
// requested length.
func Bytes(n int) []byte {
	mu64.Lock()
	defer mu64.Unlock()
	return gen64.Bytes(n)
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
