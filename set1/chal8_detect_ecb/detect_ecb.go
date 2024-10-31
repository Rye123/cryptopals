package main

import (
	"fmt"
	"os"
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

	for _, line := range lines {
		fmt.Println(line)
	}
}
