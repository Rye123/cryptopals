package main

import (
	"fmt"
	"os"

	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/util"
)

func main() {
	contents, err := os.ReadFile("./8.txt")
	if err != nil {
		fmt.Printf("File Open Error: %v\n", err)
		return
	}

	// Split by line
	var lines []string
	line := ""
	for _, c := range contents {
		if c == '\r' {
			continue
		} else if c == '\n' {
			if len(line) > 0 {
				lines = append(lines, line)
			}
			line = ""
		} else {
			line = line + string(c)
		}
	}
	if len(line) > 0 {
		lines = append(lines, line)
	}

	highestMatches := -1
	highestIdx := -1
	for i, line := range lines {
		encrypted, err := encoding.HexToBytes(line)
		if err != nil {
			fmt.Printf("Hex Decode Error: %v", err)
			return
		}

		// 1. Split into 16-byte blocks
		blockCount := (len(encrypted) / 16) + (len(encrypted) % 16)
		blocks := make([][]byte, blockCount)
		for j := 0; j < blockCount; j++ {
			blocks[j] = encrypted[j*16:(j+1)*16]
		}

		// 2. Loop through all blocks, identify number of matches
		matches := 0
		chkedIndices := make([]int, 0)
		for j, block := range blocks {
			// 2.1. If this block is already checked, ignore
			toIgnore := false
			for _, idx := range chkedIndices {
				if j == idx {
					toIgnore = true
					break
				}
			}
			if toIgnore {
				continue
			}
			chkedIndices = append(chkedIndices, j)
			
			// 2.2. Scan through blocks for matches
			for k := j + 1; k < len(blocks); k++ {
				if util.IsBytestringEqual(block, blocks[k]) {
					matches++
					chkedIndices = append(chkedIndices, k)
				}
			}
		}

		// 3. If this has more matches, set this as the index
		if matches > highestMatches {
			highestIdx = i
			highestMatches = matches
		}
	}

	fmt.Printf("The line is line %d with %d repeats.\n%s\n", highestIdx, highestMatches, lines[highestIdx])
}
