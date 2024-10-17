package encoding

import (
	"errors"
	"fmt"
	"strconv"
)

func base64Lookup(idx int) (rune, error) {
	// A-Z: 0 to 25
	if idx < 26 {
		return rune(65 + idx), nil
	}

	// a-z: 26 to 51
	if idx < 52 {
		return rune(97 + (idx - 26)), nil
	}

	// 0-9: 52 to 61
	if idx < 62 {
		return rune(48 + (idx - 52)), nil
	}

	// +: 62
	if idx == 62 {
		return '+', nil
	}

	// /: 63
	if idx == 63 {
		return '/', nil
	}

	return -1, errors.New("base64Lookup: Index exceeds 63.")
}

func BytesToBase64(bytes []byte) (string, error) {
	// 1. Convert bytestring into bits
	bits := ""
	for _, b := range bytes {
		bits += fmt.Sprintf("%08b", b)
	}

	// 2. Convert bitstring into base64
	base64Str := ""
	base64Idx := ""
	for _, c := range bits {
		base64Idx += string(c)
		if len(base64Idx) == 6 {
			idx, err := strconv.ParseInt(base64Idx, 2, 16)
			if err != nil {
				return "", err
			}
			base64Char, err := base64Lookup(int(idx))
			if err != nil {
				return "", err
			}
			base64Str += string(base64Char)
			base64Idx = ""
		}
	}

	// 3. Account for remaining data
	if len(base64Idx) > 0 {
		padCount := max(6 - len(base64Idx), 0)
		for j := 0; j < padCount; j++ {
			base64Idx += "0"
		}
		base64Idx = fmt.Sprintf("%06s", base64Idx)
		idx, err := strconv.ParseInt(base64Idx, 2, 16)
		if err != nil {
			return "", err
		}
		base64Char, err := base64Lookup(int(idx))
		if err != nil {
			return "", err
		}
		base64Str += string(base64Char)
	}

	// 4. Add padding
	remainder := len(base64Str) % 4
	padCount := min(remainder, 4 - remainder)
	for j := 0; j < padCount; j++ {
		base64Str += "="
	}
	return base64Str, nil
}

func Base64ToBytes(str string) ([]byte, error) {
	bytestr := make([]byte, 0)

	return bytestr, nil	
}
