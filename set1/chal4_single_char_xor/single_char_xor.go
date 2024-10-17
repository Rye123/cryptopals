package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/Rye123/cryptopals/lib/attacks"
	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption"
)

// A guess of a ciphertext and a key, along with the evaluated score.
type guess struct {
	hexstr string
	key byte
	score float64
}

// Returns a slice of guesses.
func evaluateHexstring(hexstr string) ([]guess, error) {
	text, err := encoding.HexToBytes(hexstr)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	// Attempt to decrypt
	guessMap, err := attacks.BreakXorSingleByte(text)
	if err != nil {
		return nil, err
	}
	
	guesses := make([]guess, 0, 256)
	for key, score := range guessMap {
		guesses = append(guesses, guess{hexstr, key, score})
	}

	return guesses, nil
}


func main() {
	f, err := os.Open("4.txt")
	if err != nil {
		fmt.Printf("File Open Error: %v\n", err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("Exit Error: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(f)
	allGuesses := make([]guess, 0, 256)
	for scanner.Scan() {
		hexstr := scanner.Text()
		guesses, err := evaluateHexstring(hexstr)
		if err != nil {
			fmt.Printf("Error evaluting string %s: %v", hexstr, err)
			continue
		}
		allGuesses = append(allGuesses, guesses...)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner Error: %v\n", err)
		return
	}

	// Sort guesses by score
	slices.SortFunc(allGuesses, func(guess1, guess2 guess) int {
		if guess1.score > guess2.score {
			return -1
		} else if guess1.score < guess2.score {
			return 1
		}
		return 0
	})

	// Print top `limit` guesses
	limit := 10
	for _, guess := range allGuesses {
		if limit == 0 {
			return
		}
		limit--
		text, err := encoding.HexToBytes(guess.hexstr)
		if err != nil {
			fmt.Printf("Hex decode error with hexstring %s: %v", guess.hexstr, err)
		}
		decrypted, err := encryption.XorSingleByte(text, guess.key)
		if err != nil {
			fmt.Printf("Decrypt error with hexstring %s, key %d: %v", guess.hexstr, guess.key, err)
		}
		fmt.Printf("%s: Key %d: %f: %s\n", guess.hexstr, guess.key, guess.score, decrypted)
	}
}
