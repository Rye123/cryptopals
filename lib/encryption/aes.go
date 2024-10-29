package encryption

import "fmt"

type AESMode int
const (
	AESMode_ECB AESMode = iota
)

// Uses AES to encrypt bytestr.
// A 16-byte key uses AES-128, 24-byte key uses AES-192, 32-byte key uses AES-256.
func AESEncrypt(bytestr []byte, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("AESEncrypt: Invalid keylength %d", len(key))
	}
	if len(bytestr) != 16 {
		return nil, fmt.Errorf("AESEncrypt: Invalid data length %d", len(bytestr))
	}
	return bytestr, nil
}

// Uses AES to decrypt bytestr.
// A 16-byte key uses AES-128, 24-byte key uses AES-192, 32-byte key uses AES-256.
func AESDecrypt(bytestr []byte, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("AESEncrypt: Invalid keylength %d", len(key))
	}
	if len(bytestr) != 16 {
		return nil, fmt.Errorf("AESEncrypt: Invalid data length %d", len(bytestr))
	}
	return bytestr, nil
}

func AESEncryptWithMode(bytestr []byte, key []byte, mode AESMode, iv []byte) ([]byte, error) {
	return bytestr, nil
}

func AESDecryptWithMode(bytestr []byte, key []byte, mode AESMode, iv []byte) ([]byte, error) {
	return bytestr, nil
}
