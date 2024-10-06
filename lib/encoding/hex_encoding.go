package encoding

import (
	"errors"
	"strconv"
)

func HexToBytes(hexstring string) ([]byte, error) {
	if len(hexstring) % 2 != 0 {
		return nil, errors.New("Expected hexstring to have even length.")
	}
	
	// Convert hexstring to bytes
	bytestr := []byte("")
	b_hex := ""
	for _, c := range hexstring {
		b_hex += string(c)
		if len(b_hex) == 2 {
			b, err := strconv.ParseInt(b_hex, 16, 16)
			if err != nil {
				return nil, err
			}
			bytestr = append(bytestr, byte(b))
			b_hex = ""
		}
	}
	
	return bytestr, nil
}

func BytesToHex(bytestring []byte) (string, error) {
	hexstring := ""
	for _, b := range bytestring {
		hexstring += strconv.FormatInt(int64(b), 16)
	}
	return hexstring, nil
}
