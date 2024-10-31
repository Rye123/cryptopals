package aes

import "fmt"

type AESMode int
const (
	AESMode_ECB AESMode = iota
)

func AESEncryptWithMode(data []byte, key []byte, mode AESMode, padding AESPadding, iv []byte) ([]byte, error) {
	blockCount := (len(data) / 16) + (len(data) % 16)
	blocks := make([][]byte, blockCount)
	
	// 1. Break data up into blocks
	for i := 0; i < blockCount; i++ {
		blocks[i] = data[i*16:(i+1)*16]
	}

	// 2. Pad last block
	block, err := padBlock(blocks[blockCount-1], padding)
	if err != nil {
		return nil, fmt.Errorf(`AESEncryptWithMode: Padding error: %v`, err)
	}
	blocks[blockCount-1] = block

	// 3. Encrypt, based on mode
	switch mode {
	case AESMode_ECB:
		for i, block := range blocks {
			result, err := aesEncrypt(block, key)
			if err != nil {
				return nil, fmt.Errorf(`AESEncryptWithMode: Encryption error: %v`, err)
			}
			blocks[i] = result
		}
	}
	
	return data, nil
}

func AESDecryptWithMode(data []byte, key []byte, mode AESMode, padding AESPadding, iv []byte) ([]byte, error) {
	// 1. Break data up into blocks

	// 2. Unpad last block

	// 3. Decrypt, based on mode
	return data, nil
}
