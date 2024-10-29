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

func circularRotateOnce(data []byte, dirLeft bool) []byte {
	result := make([]byte, len(data))
	if len(data) == 0 {
		return result
	}
	
	if dirLeft {
		leftmost := data[0]
		for i := 0; i < len(data); i++ {
			if i+1 == len(data) {
				result[i] = leftmost
			} else {
				result[i] = data[i+1]
			}
		}
	} else {
		rightmost := data[len(data)-1]
		result[0] = rightmost
		for i := 1; i < len(data); i++ {
			result[i] = data[i-1]
		}
	}
	return result
}

// Returns a new bytestring of `data` rotated by `offset` bytes. This is a circular rotation.
// Direction is determined by `dirLeft` -- if true, rotation is to the left, otherwise rotation is to the right.
func CircularRotate(data []byte, offset int, dirLeft bool) []byte {
	result := make([]byte, len(data))
	copy(result, data)
	for i := 0; i < offset; i++ {
		result = circularRotateOnce(result, dirLeft)
	}
	return result
}
