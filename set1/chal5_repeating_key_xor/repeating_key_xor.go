package main

import (
	"fmt"

	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption"
)

func main() {
	plaintext := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	encrypted, err := encryption.XorRepeating(plaintext, []byte("ICE"))
	if err != nil {
		fmt.Printf("Error when decrypting plaintext: %v\n", err);
		return
	}

	result, err := encoding.BytesToHex(encrypted)
	if err != nil {
		fmt.Printf("Error when encoding bytes: %v\n", err)
		return
	}

	fmt.Printf("Expected: %s\nResult  : %s\n", expected, result)
	if expected == result {
		fmt.Println("MATCH")
	} else {
		fmt.Println("NOT A MATCH")
	}
}
