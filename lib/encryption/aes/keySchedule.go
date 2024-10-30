package aes

import (
	"fmt"

	"github.com/Rye123/cryptopals/lib/encryption"
	"github.com/Rye123/cryptopals/lib/util"
)

func keyScheduleCore(word []byte, i int) []byte {
	result := make([]byte, len(word))
	copy(result, word)

	// 1. 8-bit left circular rotate
	result = util.CircularRotate(result, 1, true)

	// 2. Apply S-box on each byte
	for i, b := range result {
		result[i] = sBox(b)
	}

	// 3. Apply round constant
	rc := rcon(byte(i))
	result[0] = result[0] ^ rc

	return result
}

// Returns the n-byte key block for roundConstant
func genKeyBlock(roundConstant int, keySize int, prevKeyBlock []byte) ([]byte, error) {
	var err error
	keyBlock := make([]byte, 0, keySize)

	// Set initial temporary value
	t := make([]byte, 4)
	copy(t, prevKeyBlock[keySize-4:])
	t = keyScheduleCore(t, roundConstant)

	switch keySize {
	case 16:
		for i := 0; i < 4; i++ {
			t, err = encryption.XorBytes(t, prevKeyBlock[i*4 : (i+1)*4])
			if err != nil {
				return nil, err
			}
			keyBlock = append(keyBlock, t...)
		}
	case 24:
		for i := 0; i < 6; i++ {
			t, err = encryption.XorBytes(t, prevKeyBlock[i*4 : (i+1)*4])
			if err != nil {
				return nil, err
			}
			keyBlock = append(keyBlock, t...)
		}
	case 32:
		for i := 0; i < 4; i++ {
			t, err = encryption.XorBytes(t, prevKeyBlock[i*4 : (i+1)*4])
			if err != nil {
				return nil, err
			}
			keyBlock = append(keyBlock, t...)
		}

		// Apply s-box on previous value
		for i := 0; i < 4; i++ {
			t[i] = sBox(t[i])
		}
		
		for i := 0; i < 4; i++ {
			t, err = encryption.XorBytes(t, prevKeyBlock[(i+4)*4 : (i+4+1)*4])
			if err != nil {
				return nil, err
			}
			keyBlock = append(keyBlock, t...)
		}
		
	}

	return keyBlock, nil
}

func genExpandedKey(key []byte) ([]byte, error) {
	// Validate input and set expanded key size
	expKeySize := 0
	switch len(key) {
	case 16:
		expKeySize = 176
	case 24:
		expKeySize = 208
	case 32:
		expKeySize = 240
	default:
		return nil, fmt.Errorf("genExpandedKey: Invalid keylength %d", len(key))
	}
	
	// 1. Set initial len(key) bytes of the expanded key
	expKey := make([]byte, len(key), expKeySize)
	copy(expKey, key)

	prevKeyBlock := make([]byte, len(key))
	copy(prevKeyBlock, key)

	// 2. Add round keys until sufficient
	for i := 1; len(expKey) < expKeySize; i++ {
		keyBlock, err := genKeyBlock(i, len(key), prevKeyBlock)
		if err != nil {
			return nil, fmt.Errorf("genExpandedKey: Error generating round key: %v", err)
		}
		expKey = append(expKey, keyBlock...)
		copy(prevKeyBlock, keyBlock)
	}

	// 3. Truncate expanded key, if it exceeds the desired size
	expKey = expKey[:expKeySize]
	
	return expKey, nil

}
