package attacks

import (
	"errors"
	"sort"

	"github.com/Rye123/cryptopals/lib/encryption"
	"github.com/Rye123/cryptopals/lib/util"
)

// Number of blocks of the ciphertext to take to compute the average edit distance
const blockCount = 4

func guessKeylen(ciphertext []byte, keylenMin int, keylenMax int) ([]int, error) {
	type keylenGuess struct {
		keylen int
		editDist int
	}

	keylenGuesses := make([]keylenGuess, 0)
	
	// 1. Determine potential keylengths
	for l := keylenMin; l <= keylenMax; l++ {
		editDists := make([]int, 0, blockCount)
		blocks := make([][]byte, 0)
		for i := 0; i < blockCount; i++ {
			blockStart := i * l
			blockEnd := (i+1) * l
			if blockEnd >= len(ciphertext) {
				break
			}
			block := ciphertext[blockStart:blockEnd]
			blocks = append(blocks, block)
		}

		// Compute edit distances
		for i := 0; i < len(blocks); i++ {
			for j := i+1; j < len(blocks); j++ {
				editDist, err := util.HammingDistance(blocks[i], blocks[j])
				if err != nil {
					return nil, err
				}

				editDists = append(editDists, editDist)
			}
		}

		// Compute normalised edit distance
		avgEditDist := 0
		for _, editDist := range editDists {
			avgEditDist += editDist
		}
		avgEditDist /= len(editDists)
		avgEditDist /= l  // normalise avg dist by keylen

		keylenGuesses = append(keylenGuesses, keylenGuess{l, avgEditDist})
	}

	sort.Slice(keylenGuesses, func(i, j int) bool {
		return keylenGuesses[i].editDist < keylenGuesses[j].editDist
	})

	sortedKeylens := make([]int, len(keylenGuesses))
	for i, guess := range keylenGuesses {
		sortedKeylens[i] = guess.keylen
	}

	return sortedKeylens, nil
}

// Guesses the XOR key assuming a particular keylen
func guessKey(ciphertext []byte, keylen int) ([]byte, error) {
	// Break ciphertext into blocks, where each block is a monoalphabetic ciphertext
	fullKey := make([]byte, keylen)
	for idx := 0; idx < keylen; idx++ {
		// block is a monoalphabetic ciphertext
		block := make([]byte, 0)
		for i := idx; i < len(ciphertext); i += keylen {
			block = append(block, ciphertext[i])
		}

		if len(block) == 0 {
			continue
		}

		// solve the monoalphabetic ciphertext
		keyMap, err := BreakXorSingleByte(block)
		if err != nil {
			return nil, err
		}

		// sort by score and choose the highest scoring single byte key
		var key byte
		highestScore := -1.0
		for k, score := range keyMap {
			if score > highestScore {
				highestScore = score
				key = k
			}
		}
		fullKey[idx] = key
	}
		
	return fullKey, nil
}


// Returns two slices -- the first is the list of guessed keys, the second is the list of corresponding scores
func BreakXorRepeating(ciphertext []byte, keylenMin int, keylenMax int, keylenCount int) (keys [][]byte, scores []float64, err error) {
	if keylenMin > keylenMax {
		return nil, nil, errors.New("BreakXorRepeating: keylenMin > keylenMax")
	}

	// 1. Get all possible key lengths, sorted by likelihood based on edit distance
	keylens, err := guessKeylen(ciphertext, keylenMin, keylenMax)
	if err != nil {
		return nil, nil, err
	}

	// 2. Get all keys based on the guessed keylength
	for i, keylen := range keylens {
		if i >= keylenCount {
			break
		}

		key, err := guessKey(ciphertext, keylen)
		if err != nil {
			return nil, nil, err
		}

		dec, err := encryption.XorRepeating(ciphertext, key)
		if err != nil {
			return nil, nil, err
		}

		score := 1.0 - util.GetFittingQuotient(dec, util.CHAR_FREQ_ENGLISH)

		keys = append(keys, key)
		scores = append(scores, score)
	}

	return keys, scores, nil
}
