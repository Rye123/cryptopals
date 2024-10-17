package encryption

func XorRepeating(bytestr []byte, key []byte) ([]byte, error) {
	result := make([]byte, len(bytestr))
	for i, b := range bytestr {
		if len(key) > 0 {
			result[i] = b ^ key[i % len(key)]
		} else {
			result[i] = b
		}
	}
	
	return result, nil
}
