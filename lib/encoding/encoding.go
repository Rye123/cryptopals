package encoding

// Converts a hex string into base64
func HexToBase64(hexstring string) (string, error) {
	bytestr, err := HexToBytes(hexstring)
	if err != nil {
		return "", err
	}

	return BytesToBase64(bytestr)
}

// Converts base64 string into hex string
func Base64ToHex(base64string string) (string, error) {
	bytestr, err := Base64ToBytes(base64string)
	if err != nil {
		return "", err
	}

	return BytesToHex(bytestr)
}
