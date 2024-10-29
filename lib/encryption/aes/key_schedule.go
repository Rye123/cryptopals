package aes

// Returns 2^n in GF(2^8)
func rcon(in byte) byte {
	var c byte = 1

	if (in == 0) {
		return 0
	}

	for in != 1 {
		// modular multiplication in GF(2^8)
		var highBitSet byte = c & 0x80
		c <<= 1
		if highBitSet == 0x80 {
			c ^= 0x1b
		}
		
		in--
	}

	return c
}
