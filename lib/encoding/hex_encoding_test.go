package encoding

import (
	"testing"
)

func isBytestringEqual(h1, h2 []byte) bool {
	if len(h1) != len(h2) {
		return false
	}
	for i := 0; i < len(h1); i++ {
		if h1[i] != h2[i] {
			return false
		}
	}
	return true
}

func TestHexToBytes(t *testing.T) {
	// Test error
	testErr := "1fa"
	_, err := HexToBytes(testErr)
	if err == nil {
		t.Fatalf(`HexToBytes("%s") did not give an error`, testErr)
	}

	testErr = "abcdefgh"
	_, err = HexToBytes(testErr)
	if err == nil {
		t.Fatalf(`HexToBytes("%s") did not give an error`, testErr)
	}
	
	// Test strings
	tests := []string{
		"",
		"48656c6c6f20776f726c64",
		"54657374696e67206f6e652074776f207468726565",
		"1234567890",
	}

	expcs := [][]byte {
		[]byte(""),
		[]byte("Hello world"),
		[]byte("Testing one two three"),
		{0x12, 0x34, 0x56, 0x78, 0x90},
	}

	for i, test := range tests {
		result, err := HexToBytes(test)
		if err != nil {
			t.Fatalf(`HexToBytes("%s") gave an error: %v`, test, err)
		}
		
		expected := expcs[i]
		if !isBytestringEqual(result, expected) {
			t.Fatalf(`HexToBytes("%s") did not match`, test)
		}
	}
}

func TestBytesToHex(t *testing.T) {
	tests := [][]byte {
		[]byte(""),
		[]byte("Hello world"),
		[]byte("Testing one two three"),
		{0x12, 0x34, 0x56, 0x78, 0x90},
	}
	
	expcs := []string{
		"",
		"48656c6c6f20776f726c64",
		"54657374696e67206f6e652074776f207468726565",
		"1234567890",
	}

	for i, test := range tests {
		result, err := BytesToHex(test)
		if err != nil {
			t.Fatalf(`BytesToHex(tests[%d]) gave an error: %v`, i, err)
		}

		expected := expcs[i]
		if result != expected {
			t.Fatalf(`BytesToHex(tests[%d]) = "%s", expected "%s"`, i, result, expected)
		}
	}
}
