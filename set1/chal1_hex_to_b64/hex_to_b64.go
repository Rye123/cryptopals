package hex_to_b64

import (
	"fmt"
	"os"
)

// Converts a hex string into base64
func HexToBase64(hexstring string) string {
	return ""
}



func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Expected hexstrings as arguments.")
		return
	}
}
