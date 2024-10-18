package main

import (
	"fmt"
	"os"

	"github.com/Rye123/cryptopals/lib/attacks"
	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("break_repeating_key_xor base64_file")
		return
	}
	
	contents, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Printf("File Open Error: %v\n", err)
		return
	}

	// Remove newlines from contents
	encoded := ""
	for _, c := range contents {
		if c == '\n' {
			continue
		}
		encoded += string(c)
	}

	encrypted, err := encoding.Base64ToBytes(encoded)
	if err != nil {
		fmt.Printf("Expected a file in Base64: %v\n", err)
		return
	}

	keys, scores, err := attacks.BreakXorRepeating(encrypted, 2, 40, 3)
	if err != nil {
		fmt.Printf("BreakXorRepeating Error: %v\n", err)
		return
	}

	// Identify key with highest score
	var brokenKey []byte
	highestScore := -1.0
	for i, key := range keys {
		if scores[i] > highestScore {
			highestScore = scores[i]
			brokenKey = key
		}
	}
		
	keyHex, err := encoding.BytesToHex(brokenKey)
	if err != nil {
		fmt.Printf("Warning: Could not encode key to hex, key: %s\n", brokenKey)
		return
	}
	decrypted, err := encryption.XorRepeating(encrypted, brokenKey)
	if err != nil {
		fmt.Printf("Warning: Could not decrypt ciphertext with key %s\n", keyHex)
		return
	}

	fmt.Printf("Key detected: 0x%s, Score: %f\n---\n%s\n", keyHex, highestScore, decrypted)
}
