package util

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
	return 0, nil
}
