package main

import (
	"fmt"
	"sort"
	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption"
	"github.com/Rye123/cryptopals/lib/attacks"
)

// Prints the top five possible decryptions of the given string
func evaluateHexstring(hexstr string) {
	text, err := encoding.HexToBytes(hexstr)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Attempt to decrypt
	scores := make(map[int]float64)
	for i := 0; i < 256; i++ {
		decrypted, err := encryption.XorSingleByte(text, byte(i))
		if err != nil {
			fmt.Printf("Error decrypting with key %d: %v\n", i, err)
			scores[i] = 0.0
			continue
		}

		scores[i] = attacks.ScoreText(decrypted)
	}

	// Sort scores by value
	keys := make([]int, 0, len(scores))
	for key := range scores {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return scores[keys[i]] > scores[keys[j]]
	})

	// Print top 3 keys and decrypted values
	count := 3
	fmt.Printf("String: %s:\n", hexstr)
	for _, key := range keys {
		if count == 0 {
			return
		}
		count--
		decrypted, err := encryption.XorSingleByte(text, byte(key))
		if err != nil {
			decrypted = []byte("(Invalid string)")
		}
		fmt.Printf("Key %d: %f: %s\n", key, scores[key], decrypted)
	}
	fmt.Printf("\n")
}

func main() {
	test_str := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	evaluateHexstring(test_str)
}
