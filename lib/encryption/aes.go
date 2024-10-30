package encryption

type AESMode int
const (
	AESMode_ECB AESMode = iota
)

func AESEncryptWithMode(bytestr []byte, key []byte, mode AESMode, iv []byte) ([]byte, error) {
	return bytestr, nil
}

func AESDecryptWithMode(bytestr []byte, key []byte, mode AESMode, iv []byte) ([]byte, error) {
	return bytestr, nil
}
