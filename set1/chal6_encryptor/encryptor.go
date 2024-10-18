package main

import (
	"fmt"
	"os"

	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("encryptor hexkey file")
		return
	}
	keyHex := args[1]
	filename := args[2]

	plaintext, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error when reading file %s: %v\n", filename, err)
		return
	}

	key, err := encoding.HexToBytes(keyHex)
	if err != nil {
		fmt.Printf("Error when decoding key %s: %v\n", keyHex, err)
		return
	}

	encrypted, err := encryption.XorRepeating(plaintext, key)
	if err != nil {
		fmt.Printf("Error encrypting file %s with key %s: %v\n", filename, keyHex, err)
		return
	}

	encoded, err := encoding.BytesToBase64(encrypted)

	filenameEnc := filename + ".enc"
	outfile, err := os.Create(filenameEnc)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filenameEnc, err)
		return
	}

	_, err = outfile.WriteString(encoded)
	if err != nil {
		fmt.Printf("Error writing encrypted bytes to file %s: %v\n", filenameEnc, err)
		return
	}
	fmt.Printf("Wrote encrypted file: %s\n", filenameEnc)

	err = outfile.Close()
	if err != nil {
		fmt.Printf("Error closing file %s: %v\n", filenameEnc, err)
	}
}
