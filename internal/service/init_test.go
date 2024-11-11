package service_test

import (
	"crypto/rsa"
	"math/rand"
)

const testSeed = 1

func GenerateRSAKeys(bits int) (*rsa.PrivateKey, error) {
	rnd := rand.New(rand.NewSource(testSeed)) //nolint - this is weak random number generator
	return rsa.GenerateKey(rnd, bits)
}
