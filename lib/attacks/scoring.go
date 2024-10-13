package attacks

import "math"

var CHAR_FREQ_ENGLISH = map[rune]float64{
	'A': 0.082,
	'B': 0.015,
	'C': 0.028,
	'D': 0.043,
	'E': 0.127,
	'F': 0.022,
	'G': 0.020,
	'H': 0.061,
	'I': 0.070,
	'J': 0.0015,
	'K': 0.0077,
	'L': 0.04,
	'M': 0.024,
	'N': 0.067,
	'O': 0.075,
	'P': 0.019,
	'Q': 0.00095,
	'R': 0.06,
	'S': 0.063,
	'T': 0.091,
	'U': 0.028,
	'V': 0.0098,
	'W': 0.024,
	'X': 0.0015,
	'Y': 0.020,
	'Z': 0.00074,
}

// Returns if the given byte represents a printable character in ASCII (i.e. alphabet, number, punctuation)
func isPrintable(b byte) bool {
	return b >= 0x20 && b <= 0x7E
}

// Scores how likely the text is to be English
func ScoreText(text []byte) float64 {
	// If text contains non-printable characters, probably not English
	textFreq := map[rune]float64{
		'A': 0.0,
		'B': 0.0,
		'C': 0.0,
		'D': 0.0,
		'E': 0.0,
		'F': 0.0,
		'G': 0.0,
		'H': 0.0,
		'I': 0.0,
		'J': 0.0,
		'K': 0.0,
		'L': 0.0,
		'M': 0.0,
		'N': 0.0,
		'O': 0.0,
		'P': 0.0,
		'Q': 0.0,
		'R': 0.0,
		'S': 0.0,
		'T': 0.0,
		'U': 0.0,
		'V': 0.0,
		'W': 0.0,
		'X': 0.0,
		'Y': 0.0,
		'Z': 0.0,		
	}

	letterCount := 0
	for _, b := range text {
		if !isPrintable(b) {
			return 0.0
		}

		// Add to textFreq if letter
		if b >= 0x41 && b <= 0x5A {
			letterCount++
			textFreq[rune(b)]++
		}
		if b >= 0x61 && b <= 0x7A {
			letterCount++
			textFreq[rune(b - 32)]++
		}
	}

	// Evaluate frequencies
	if letterCount > 0 {
		for c, _ := range textFreq {
			textFreq[c] /= float64(letterCount)
		}
	}
	
	// Compute mean squared deviation between actual and english frequency
	msd := 0.0
	for c, _ := range textFreq {
		deviation := textFreq[c] - CHAR_FREQ_ENGLISH[c]
		msd += (deviation * deviation)
	}
	
	// Compute Shannon entropy
	entropy := 0.0
	for c, _ := range textFreq {
		if textFreq[c] != 0 {
			entropy -= textFreq[c] * math.Log(textFreq[c])
		}
	}

	// Compute final score
	score := entropy - msd
	
	return score
}
