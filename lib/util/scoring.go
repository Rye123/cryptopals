package util

import (
	"math"
	"unicode"
)

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
	' ': 0.254,
}

// Returns the fitting quotient of the text from a reference frequency
func GetFittingQuotient(text []byte, refFreq map[rune]float64) float64 {
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
		' ': 0.0,
	}

	for _, b := range text {
		// Add to textFreq if letter
		if b >= 0x41 && b <= 0x5A {
			textFreq[rune(b)]++
		}
		if b >= 0x61 && b <= 0x7A {
			textFreq[rune(b - 32)]++
		}
		if b == 0x20 {
			textFreq[rune(b)]++
		}
	}

	// Evaluate frequencies
	if len(text) > 0 {
		for c, _ := range textFreq {
			textFreq[c] /= float64(len(text))
		}
	}
	
	// Compute fitting quotient
	fq := 0.0
	for c, _ := range textFreq {
		deviation := textFreq[c] - refFreq[c]
		fq += math.Abs(deviation)
	}
	fq /= float64(len(refFreq))

	return fq
	
}

// Returns the Shannon entropy of a text
func GetShannonEntropy(text []byte) float64 {
	if len(text) == 0 {
		return 0.0
	}
	
	freq := make(map[byte]float64)

	// Get count
	for _, b := range text {
		b := byte(unicode.ToUpper(rune(b)))
		freq[b]++
	}

	// Get freq
	for _, b := range text {
		b := byte(unicode.ToUpper(rune(b)))
		freq[b] = freq[b] / float64(len(text))
	}

	entropy := 0.0
	for _, char_freq := range freq {
		if char_freq != 0 {
			entropy -= char_freq * math.Log(char_freq)
		}
	}

	return entropy
}
