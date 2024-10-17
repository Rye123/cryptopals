package encryption

import (
	"testing"
	"github.com/Rye123/cryptopals/lib/util"
)

func TestXorSingleByte(t *testing.T) {
	tests_bytestr := [][]byte{
		[]byte(""),
		[]byte("The quick brown fox"),
		{0x1b, 0x37, 0x37, 0x33, 0x31, 0x36, 0x3f, 0x78, 0x15, 0x1b, 0x7f, 0x2b, 0x78, 0x34, 0x31, 0x33, 0x3d, 0x78, 0x39, 0x78, 0x28, 0x37, 0x2d, 0x36, 0x3c, 0x78, 0x37, 0x3e, 0x78, 0x3a, 0x39, 0x3b, 0x37, 0x36},
	}

	tests_key := []byte{
		'x',
		0x69,
		0x58,
	}

	expcs := [][]byte{
		[]byte(""),
		{0x3d, 0x01, 0x0c, 0x49, 0x18, 0x1c, 0x00, 0x0a, 0x02, 0x49, 0x0b, 0x1b, 0x06, 0x1e, 0x07, 0x49, 0x0f, 0x06, 0x11},
		[]byte("Cooking MC's like a pound of bacon"),
	}

	if len(tests_bytestr) != len(tests_key) || len(tests_key) != len(expcs) {
		t.Fatalf(`TestXorSingleByte: Test error: tests count do not match`)
	}

	for i := 0; i < len(expcs); i++ {
		result, err := XorSingleByte(tests_bytestr[i], tests_key[i])
		if err != nil {
			t.Fatalf(`TestXorSingleByte(tests_bytestr[%d], tests_key[%d]) gave error: %v`, i, i, err)
		}
		if !util.IsBytestringEqual(result, expcs[i]) {
			t.Fatalf(`TestXorSingleByte(tests_bytestr[%d], tests_key[%d]) != expcs[%d]`, i, i, i)
		}
	}
}
