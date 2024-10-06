package main

import (
	"os"
	"fmt"
	"github.com/Rye123/cryptopals/lib/encoding"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		return
	}

	for _, arg := range args[1:] {
		result, err := encoding.HexToBase64(arg)
		if err != nil {
			fmt.Printf("%s: Error: %v\n", arg, err)
		} else {
			fmt.Printf("%s: %s\n", arg, result)
		}
	}
}
