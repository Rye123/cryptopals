package main

import (
	"fmt"
	"os"

	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption/aes"
)

func main() {
	key := []byte("YELLOW SUBMARINE")
	contents, err := os.ReadFile("./7.txt")
	if err != nil {
		fmt.Printf("File Open Error: %v\n", err)
		return
	}

	// Remove newlines from contents
	encoded := ""
	for _, c := range contents {
		if c == '\n' || c == '\r' {
			continue
		}
		encoded += string(c)
	}

	encrypted, err := encoding.Base64ToBytes(encoded)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return
	}
	
	decrypted, err := aes.AESDecrypt(encrypted, key, aes.AESMode_ECB, aes.AESPadding_EMPTY, nil)
	if err != nil {
		fmt.Printf("AES Decryption Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %s\n", decrypted)
	
}
