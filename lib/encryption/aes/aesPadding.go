package aes

import "errors"

type AESPadding int
const (
	AESPadding_EMPTY AESPadding = iota
)

// Returns a block padded to 16 bytes, or an error if the block exceeds
// 16 bytes.
func getPaddedBlock(bytestr []byte, padType AESPadding) ([]byte, error) {
	if len(bytestr) > 16 {
		return nil, errors.New("expected block to be less than 16 bytes.")
	}

	result := make([]byte, 0, 16)
	result = append(result, bytestr...)

	for len(result) < 16 {
		switch padType {
		case AESPadding_EMPTY:
			result = append(result, byte(0))
		default:
			return nil, errors.New("unknown AES padding type.")
		}
	}
	return result, nil
}
