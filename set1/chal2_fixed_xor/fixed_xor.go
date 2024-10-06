package main

import (
	"fmt"
	"os"

	"github.com/Rye123/cryptopals/lib/encoding"
	"github.com/Rye123/cryptopals/lib/encryption"
)


func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Errorf("fixed_xor hexstring1 hexstring2\nReturns the XORed result in hex.")
		return
	}

	b1, err := encoding.HexToBytes(args[1])
	if err != nil {
		fmt.Errorf("hexstring1 error: %v\n", err)
		return
	}

	b2, err := encoding.HexToBytes(args[2])
	if err != nil {
		fmt.Errorf("hexstring2 error: %v\n", err)
		return
	}

	result_b, err := encryption.XorBytes(b1, b2)
	if err != nil {
		fmt.Errorf("fixed_xor error: %v\n", err)
		return
	}

	result_hex, err := encoding.BytesToHex(result_b)
	if err != nil {
		fmt.Errorf("fixed_xor error: %v\n", err)
		return
	}
	fmt.Printf("%s\n", result_hex)
}
