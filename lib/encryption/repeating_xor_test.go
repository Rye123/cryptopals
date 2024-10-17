package encryption

import (
	"testing"
	"github.com/Rye123/cryptopals/lib/util"
)

func TestXorRepeating(t *testing.T) {
	tests_bytestr := [][]byte{
		[]byte(""),
		[]byte("The quick brown fox"),
		[]byte("lorem ipsum dolor SIT AMET"),
		[]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"),
	}

	tests_key := [][]byte{
		[]byte("password"),
		[]byte("password"),
		{0xde, 0xad, 0xbe, 0xef},
		[]byte("ICE"),
	}

	expcs := [][]byte{
		[]byte(""),
		{0x24, 0x09, 0x16, 0x53, 0x06, 0x1a, 0x1b, 0x07, 0x1b, 0x41, 0x11, 0x01, 0x18, 0x18, 0x1c, 0x44, 0x16, 0x0e, 0x0b},
		{0xb2, 0xc2, 0xcc, 0x8a, 0xb3, 0x8d, 0xd7, 0x9f, 0xad, 0xd8, 0xd3, 0xcf, 0xba, 0xc2, 0xd2, 0x80, 0xac, 0x8d, 0xed, 0xa6, 0x8a, 0x8d, 0xff, 0xa2, 0x9b, 0xf9},
		{0x0b, 0x36, 0x37, 0x27, 0x2a, 0x2b, 0x2e, 0x63, 0x62, 0x2c, 0x2e, 0x69, 0x69, 0x2a, 0x23, 0x69, 0x3a, 0x2a, 0x3c, 0x63, 0x24, 0x20, 0x2d, 0x62, 0x3d, 0x63, 0x34, 0x3c, 0x2a, 0x26, 0x22, 0x63, 0x24, 0x27, 0x27, 0x65, 0x27, 0x2a, 0x28, 0x2b, 0x2f, 0x20, 0x43, 0x0a, 0x65, 0x2e, 0x2c, 0x65, 0x2a, 0x31, 0x24, 0x33, 0x3a, 0x65, 0x3e, 0x2b, 0x20, 0x27, 0x63, 0x0c, 0x69, 0x2b, 0x20, 0x28, 0x31, 0x65, 0x28, 0x63, 0x26, 0x30, 0x2e, 0x27, 0x28, 0x2f},
	}

	if len(tests_bytestr) != len(tests_key) || len(tests_key) != len(expcs) {
		t.Fatalf(`TestXorRepeating: Test error: tests count do not match`)
	}

	for i := 0; i < len(expcs); i++ {
		result, err := XorRepeating(tests_bytestr[i], tests_key[i])
		if err != nil {
			t.Fatalf(`TestXorRepeating(tests_bytestr[%d], tests_key[%d]) gave error: %v`, i, i, err)
		}
		if !util.IsBytestringEqual(result, expcs[i]) {
			t.Fatalf(`TestXorRepeating(tests_bytestr[%d], tests_key[%d]) != expcs[%d]`, i, i, i)
		}
	}
	
}
