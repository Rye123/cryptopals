package encryption

import "errors"

func XorBytes(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, errors.New("lengths of bytestrings do not match")
	}

	result := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		result[i] = b1[i] ^ b2[i]
	}

	return result, nil
}
