package aes

import (
	"fmt"

	_ "github.com/Rye123/cryptopals/lib/encryption"
	"github.com/Rye123/cryptopals/lib/util"
)

func keyScheduleCore(word []byte, i int) []byte {
	result := make([]byte, len(word))
	copy(result, word)

	// 1. 8-bit left circular rotate
	result = util.CircularRotate(result, 2, true)

	// 2. Apply S-box on each byte
	for i, b := range result {
		result[i] = sBox(b)
	}

	// 3. Apply round constant
	rc := rcon(byte(i))
	result[0] = result[0] ^ rc

	return result
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
	expKeySize = expKeySize + 1 // to ignore error

	return nil, nil
	/*// 1. Set initial len(key) bytes of the expanded key
	expKey := make([]byte, len(key), expKeySize)
	copy(expKey, key)
	i := 1

	// Initialise temp variable for initial 16 bytes
	t := make([]byte, 4)
	copy(t, expKey[0:4])
	t = keyScheduleCore(t, i)

	// 2. Set next 
	for i < 4 {
		// XOR t with 4-byte block n bytes before, where n is length of the key
		blockIdx := len(expKey) - len(key)
		block := expKey[blockIdx:blockIdx+4]
		value, err := encryption.XorBytes(block, t)
		if err != nil {
			return nil, fmt.Errorf("genExpandedKey: Error when XORing bytes: %v", err)
		}
		copy(expKey[i*4:(i+1)*4], value)
		
	}*/

}
