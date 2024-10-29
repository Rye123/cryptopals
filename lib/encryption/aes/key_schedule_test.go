package aes

import (
	"testing"
)

func TestRcon(t *testing.T) {
	tests := []byte{
		0x01, 0x02, 0x03, 0x04, 0x05,
		0x06, 0x07, 0x08, 0x09, 0x0a,
	}
	expcs := []byte{
		0x01, 0x02, 0x04, 0x08, 0x10,
		0x20, 0x40, 0x80, 0x1b, 0x36,
	}
	if len(tests) != len(expcs) {
		t.Fatalf("TestRcon: Size of test and expc slices don't match")
	}
	for i := 0; i < len(tests); i++ {
		result := rcon(tests[i])
		if result != expcs[i] {
			t.Fatalf("TestRcon: results[%d] = %x, expected %x", i, result, expcs[i])
		}
	}
}
		
