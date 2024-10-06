package encryption

func XorSingleByte(bytestr []byte, key byte) ([]byte, error) {
	result := make([]byte, len(bytestr))
	for i, b := range bytestr {
		result[i] = b ^ key
	}
	return result, nil
}
