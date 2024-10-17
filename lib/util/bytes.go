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
