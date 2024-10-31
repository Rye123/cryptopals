package aes

type AESMode int
const (
	AESMode_ECB AESMode = iota
)

func AESEncryptWithMode(data []byte, key []byte, mode AESMode, padding AESPadding, iv []byte) ([]byte, error) {
	// 1. Break data up into blocks

	// 2. Pad last block

	// 3. Encrypt, based on mode
	
	return data, nil
}

func AESDecryptWithMode(data []byte, key []byte, mode AESMode, padding AESPadding, iv []byte) ([]byte, error) {
	// 1. Break data up into blocks

	// 2. Unpad last block

	// 3. Decrypt, based on mode
	return data, nil
}
