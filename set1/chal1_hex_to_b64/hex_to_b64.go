package hex_to_b64

import (
	"errors"
	"fmt"
	"os"
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

// Converts a hex string into base64
func HexToBase64(hexstring string) (string, error) {
	if len(hexstring) % 2 != 0 {
		return "", errors.New("Expected hexstring to have even length.")
	}
	
	// 1. Convert hexstring to bytes
	bytestr := []byte("")
	b_hex := ""
	for _, c := range hexstring {
		b_hex += string(c)
		if len(b_hex) == 2 {
			b, err := strconv.ParseInt(b_hex, 16, 16)
			if err != nil {
				return "", err
			}
			bytestr = append(bytestr, byte(b))
			b_hex = ""
		}
		
	}

	return BytesToBase64(bytestr)
}



func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Expected hexstrings as arguments.")
		return
	}
}
