# randchars
[![GoDoc](https://godoc.org/github.com/mohae/randchars?status.svg)](https://godoc.org/github.com/mohae/randchars)[![Build Status](https://travis-ci.org/mohae/randchars.png)](https://travis-ci.org/mohae/randchars)

Fast random ASCII character generation.  There are both PRNG and CSPRNG based implementations.  For convenience, package level globals are provided.  These package globals are safe for concurrent use.

## PRNG
To use the PRNG based random character generation:

    import "github.com/mohae/randchars"

The PRNGs are seeded using a random value obtained from `crypto/rand`.  The PRNGs can be re-seeded, either using a user supplied value or with a value obtained from `crypto/rand`.

### Generator
Generator quickly generates random characters of an arbitrary length with the following character set options: `a-zA-Z0-9`, `a-zA-Z`, `a-z0-9`, `a-z`, `A-Z0-9`, `A-Z`, and [Base64](https://tools.ietf.org/html/rfc4648).

Generator uses a RNG that implements [PCG](http://www.pcg-random.org) written by Damian Gryski: [go-pcgr](https://github.com/dgryski/go-pcgr)

This fulfills the `Generatorer` interface.

### Base64Generator
The Base64Generator generates random characters of an arbitrary length using the base 64 alphabet as shown in [Table 1 of RFC 4648](https://tools.ietf.org/html/rfc4648) and uses a PRNG that implements [XORoShiRo128+](http://xoroshiro.di.unimi.it/) written by Damian Gryski: [go-xoroshiro](https://github.com/dgryski/go-xoroshiro). This generator is slightly faster than using `Generator.Base64()` and existed before `Generator` had a `Base64` method, which was added to `Generator` so it could fulfill the `Generatorer` interface.

## CSPRNG
For use-cases that require a CSPRNG, a CSPRNG based implementation is provided.

To use the CSPRNG based random character generation:

    import "github.com/mohae/randchars/crandchars"

This version uses the stdlib's `crypto/rand` package.  The `Generator` caches a number of random bytes.  The cache is refilled whenever it is exhausted.  This speeds up the process of generating random characters.  If a local `Generator` is being used, the cache size can be specified by using the `NewGenerator()` func.

For convenience, a thread-safe package level `Generator` is provided.

The CSPRNg Generator implements the `randchars.Generatorer` interface.

## License
MIT Licensed.  See the LICENSE file.
