package aes

import "fmt"

// Uses AES to encrypt bytestr.
// A 16-byte key uses AES-128, 24-byte key uses AES-192, 32-byte key uses AES-256.
func aesEncrypt(bytestr []byte, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("AESEncrypt: Invalid keylength %d", len(key))
	}
	if len(bytestr) != 16 {
		return nil, fmt.Errorf("AESEncrypt: Invalid data length %d", len(bytestr))
	}

	// Generate expanded key
	expKey, err := genExpandedKey(key)
	if err != nil {
		return nil, fmt.Errorf("AESEncrypt: Key expansion error: %v", err)
	}
	roundCount := len(expKey) / 16

	// Set up state array:
	// b00 b04 b08 b12
	// b01 b05 b09 b13
	// b02 b06 b10 b14
	// b03 b07 b11 b15
	state := []byte{
		bytestr[0], bytestr[4], bytestr[8], bytestr[12],
		bytestr[1], bytestr[5], bytestr[9], bytestr[13],
		bytestr[2], bytestr[6], bytestr[10], bytestr[14],
		bytestr[3], bytestr[7], bytestr[11], bytestr[15],
	}

	// Round 1: AddRoundKey
	roundKey := expKey[0:16]
	state, err = addRoundKey(state, roundKey)
	if err != nil {
		return nil, fmt.Errorf("AESEncrypt: Add round key error: %v", err)
	}

	// Round i:
	for i := 1; i < roundCount - 1; i++ {
		roundKey = expKey[16*i:16*(i+1)]
		state = subBytes(state)
		state = rowShift(state)
		state = mixCols(state)
		state, err = addRoundKey(state, roundKey)
		if err != nil {
			return nil, fmt.Errorf("AESEncrypt: Add round key error: %v", err)
		}
	}

	// Final Round:
	roundKey = expKey[len(expKey)-16:]
	state = subBytes(state)
	state = rowShift(state)
	state, err = addRoundKey(state, roundKey)
	if err != nil {
		return nil, fmt.Errorf("AESEncrypt: Add round key error: %v", err)
	}

	// Generate encrypted block
	block := []byte{
		state[0], state[4], state[8], state[12],
		state[1], state[5], state[9], state[13],
		state[2], state[6], state[10], state[14],
		state[3], state[7], state[11], state[15],
	}
		
	return block, nil
}

// Uses AES to decrypt bytestr.
// A 16-byte key uses AES-128, 24-byte key uses AES-192, 32-byte key uses AES-256.
func aesDecrypt(bytestr []byte, key []byte) ([]byte, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, fmt.Errorf("AESDecrypt: Invalid keylength %d", len(key))
	}
	if len(bytestr) != 16 {
		return nil, fmt.Errorf("AESDecrypt: Invalid data length %d", len(bytestr))
	}

	// Generate expanded key
	expKey, err := genExpandedKey(key)
	if err != nil {
		return nil, fmt.Errorf("AESDecrypt: Key expansion error: %v", err)
	}
	roundCount := len(expKey) / 16

	// Set up state array:
	// b00 b04 b08 b12
	// b01 b05 b09 b13
	// b02 b06 b10 b14
	// b03 b07 b11 b15
	state := []byte{
		bytestr[0], bytestr[4], bytestr[8], bytestr[12],
		bytestr[1], bytestr[5], bytestr[9], bytestr[13],
		bytestr[2], bytestr[6], bytestr[10], bytestr[14],
		bytestr[3], bytestr[7], bytestr[11], bytestr[15],
	}

	// Round 1: AddRoundKey
	roundKey := expKey[len(expKey)-16:]
	state, err = addRoundKey(state, roundKey)
	if err != nil {
		return nil, fmt.Errorf("AESEncrypt: Add round key error: %v", err)
	}

	// Round i:
	for i := roundCount - 2; i > 0; i-- {
		roundKey = expKey[16*i:16*(i+1)]
		state = rowShiftInv(state)
		state = subBytesInv(state)
		state, err = addRoundKey(state, roundKey)
		if err != nil {
			return nil, fmt.Errorf("AESDecrypt: Add round key error: %v", err)
		}
		state = mixColsInv(state)
	}

	// Final Round:
	roundKey = expKey[0:16]
	state = rowShiftInv(state)
	state = subBytesInv(state)
	state, err = addRoundKey(state, roundKey)
	if err != nil {
		return nil, fmt.Errorf("AESEncrypt: Add round key error: %v", err)
	}

	// Generate encrypted block
	block := []byte{
		state[0], state[4], state[8], state[12],
		state[1], state[5], state[9], state[13],
		state[2], state[6], state[10], state[14],
		state[3], state[7], state[11], state[15],
	}
		
	return block, nil
}
