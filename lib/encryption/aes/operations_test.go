package aes

import (
	"testing"

	"github.com/Rye123/cryptopals/lib/util"
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
		
func TestSBox(t *testing.T) {
	tests := []byte{
		0x9a,
		0x01,
		0xff,
		0xde,
		0xad,
		0xbe,
		0xef,
	}
	expcs := []byte{
		0xb8,
		0x7c,
		0x16,
		0x1d,
		0x95,
		0xae,
		0xdf,
	}
	if len(tests) != len(expcs) {
		t.Fatalf("TestSBox: Size of test and expc slices don't match")
	}

	// Test SBox
	for i := 0; i < len(tests); i++ {
		result := sBox(tests[i])
		if result != expcs[i] {
			t.Fatalf("TestSBox: results[%d] = %x, expected %x", i, result, expcs[i])
		}
	}

	// Test SBoxInv
	for i := 0; i < len(tests); i++ {
		result := sBoxInv(expcs[i])
		if result != tests[i] {
			t.Fatalf("TestSBox (Inv): results[%d] = %x, expected %x", i, result, tests[i])
		}
	}
}

func TestRowShift(t *testing.T) {
	state := []byte{
		0xde, 0xad, 0xbe, 0xef,
		0xca, 0xfe, 0xba, 0xbe,
		0xaa, 0xbb, 0xcc, 0xdd,
		0x01, 0x02, 0x03, 0x04,
	}
	expcNewState := []byte{
		0xde, 0xad, 0xbe, 0xef,
		0xfe, 0xba, 0xbe, 0xca,
		0xcc, 0xdd, 0xaa, 0xbb,
		0x04, 0x01, 0x02, 0x03,
	}
	newState := rowShift(state)
	origState := rowShiftInv(expcNewState)

	if !util.IsBytestringEqual(newState, expcNewState) {
		t.Fatalf("TestRowShift: newState != state. newState = %s", util.BytestringAsString(newState))
	}
	if !util.IsBytestringEqual(origState, state) {
		t.Fatalf("TestRowShift (Inv): origState != state. origState = %s", util.BytestringAsString(origState))
	}
}
