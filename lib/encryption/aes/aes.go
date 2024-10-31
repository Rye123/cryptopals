package aes

import "fmt"

type AESMode int
const (
	AESMode_ECB AESMode = iota
)

func AESEncrypt(data []byte, key []byte, mode AESMode, padding AESPadding, iv []byte) ([]byte, error) {
	blockCount := (len(data) / 16) + (len(data) % 16)
	blocks := make([][]byte, blockCount)
	
	// 1. Break data up into blocks
	for i := 0; i < blockCount; i++ {
		blocks[i] = data[i*16:(i+1)*16]
	}

	// 2. Pad last block
	block, err := padBlock(blocks[blockCount-1], padding)
	if err != nil {
		return nil, fmt.Errorf(`AESEncrypt: Padding error: %v`, err)
	}
	blocks[blockCount-1] = block

	// 3. Encrypt, based on mode
	switch mode {
	case AESMode_ECB:
		for i, block := range blocks {
			result, err := aesEncrypt(block, key)
			if err != nil {
				return nil, fmt.Errorf(`AESEncrypt: Encryption error: %v`, err)
			}
			blocks[i] = result
		}
	}

	// 4. Combine encrypted blocks
	encrypted := make([]byte, 0, blockCount * 16)
	for i := 0; i < blockCount; i++ {
		encrypted = append(encrypted, blocks[i]...)
	}
	
	return encrypted, nil
}

func AESDecrypt(data []byte, key []byte, mode AESMode, padding AESPadding, iv []byte) ([]byte, error) {
	blockCount := (len(data) / 16) + (len(data) % 16)
	blocks := make([][]byte, blockCount)
	
	// 1. Break data up into blocks
	for i := 0; i < blockCount; i++ {
		blocks[i] = data[i*16:(i+1)*16]
	}

	// 2. Unpad last block
	block, err := unpadBlock(blocks[blockCount-1], padding)
	if err != nil {
		return nil, fmt.Errorf(`AESDecrypt: Padding error: %v`, err)
	}
	blocks[blockCount-1] = block

	// 3. Decrypt, based on mode
	switch mode {
	case AESMode_ECB:
		for i, block := range blocks {
			result, err := aesDecrypt(block, key)
			if err != nil {
				return nil, fmt.Errorf(`AESDecrypt: Decryption error: %v`, err)
			}
			blocks[i] = result
		}
	}

	// 4. Combine decrypted blocks
	decrypted := make([]byte, 0, blockCount * 16)
	for i := 0; i < blockCount; i++ {
		decrypted = append(decrypted, blocks[i]...)
	}
	
	return decrypted, nil
}
