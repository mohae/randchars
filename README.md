# randchars
[![GoDoc](https://godoc.org/github.com/mohae/randchars?status.svg)](https://godoc.org/github.com/mohae/randchars)[![Build Status](https://travis-ci.org/mohae/randchars.png)](https://travis-ci.org/mohae/randchars)

Fast, `Generator`, and faster, `Gen64`, random ASCII character generation.  For convenience, package level globals are provided.  These package globals are safe for concurrent use.  By default, the seed is provided by `crypto\rand`.

## PRNG
To use the PRNG based random character generation:

    import "github.com/mohae/randchars"

### Generator
Generator quickly generates random characters of an arbitrary length with the following character set options: `a-zA-Z0-9`, `a-zA-Z`, `a-z0-9`, `a-z`, `A-Z0-9`, and `A-Z`.

Generator uses a RNG that implements [PCG](http://www.pcg-random.org) written by Damian Gryski: [go-pcgr](https://github.com/dgryski/go-pcgr)

### Gen64
Gen64 generates random characters of an arbitrary length using a 64 character set: `a-z`, `A-Z`, `0-9`, `=`, and `_`.  This is ~20% faster than using Generator.

Gen64 uses a RNG that implements [XORoShiRo128+](http://xoroshiro.di.unimi.it/) written by Damian Gryski: [go-xoroshiro](https://github.com/dgryski/go-xoroshiro)


## CSPRNG
For use-cases that require a CSPRNG, a CSPRNG based implementation is provided.

To use the CSPRNG based random character generation:

    import "github.com/mohae/randchars/crandchars"

This version uses the stdlib's `crypto/rand` package.  The `Generator` caches a number of random bytes.  The cache is refilled whenever it is exhausted.  If a local `Generator` is being used, the cache size can be specified by using the `NewGenerator()` func.

For convenience, a thread-safe package level `Generator` is provided.

## License
MIT Licensed.  See the LICENSE file.
