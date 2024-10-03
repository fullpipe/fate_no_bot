package main

import (
	"crypto/rand"
	"math/big"
)

// RandomInt returns a uniform random value in [0, max). It panics if max <= 0.
func RandomInt(max int) int {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}

	return int(num.Int64())
}
