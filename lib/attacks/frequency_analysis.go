package attacks

import (
	"github.com/Rye123/cryptopals/lib/encryption"
	"github.com/Rye123/cryptopals/lib/util"
)

// Returns a map that maps each single-byte key to the score.
func BreakXorSingleByte(ciphertext []byte) (map[byte]float64, error) {
	guesses := make(map[byte]float64)
	
	for i := 0; i < 256; i++ {
		key := byte(i)
		dec, err := encryption.XorSingleByte(ciphertext, key)
		if err != nil {
			return nil, err
		}
		
		guesses[key] = 1.0 - util.GetFittingQuotient(dec, util.CHAR_FREQ_ENGLISH)
	}
	return guesses, nil
}
