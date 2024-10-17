package util

import (
	"errors"
	"math/bits"
)

func IsBytestringEqual(h1, h2 []byte) bool {
	if len(h1) != len(h2) {
		return false
	}
	for i := 0; i < len(h1); i++ {
		if h1[i] != h2[i] {
			return false
		}
	}
	return true
}

// Returns the hamming distance between two bytestrings
func HammingDistance(h1, h2 []byte) (int, error) {
	if len(h1) != len(h2) {
		return -1, errors.New("HammingDistance: expected equal length strings")
	}

	hammingDist := 0
	for i := 0; i < len(h1); i++ {
		xorVal := uint(h1[i] ^ h2[i])
		hammingDist += bits.OnesCount(xorVal)
	}
	
	return hammingDist, nil
}
