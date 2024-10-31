package aes

import (
	"testing"

	"github.com/Rye123/cryptopals/lib/util"
)

func TestGetPaddedBlock(t *testing.T) {
	// Test invalid length
	invalid := []byte("AAAABBBBCCCCDDDDe")
	_, err := getPaddedBlock(invalid, AESPadding_EMPTY)
	if err == nil {
		t.Fatalf("TestGetPaddedBlock: Expected error from invalid bytestring of length %d", len(invalid))
	}

	// Test equal length
	equal := []byte("AAAABBBBCCCCDDDD")
	result, err := getPaddedBlock(equal, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestGetPaddedBlock: Error when processing bytestring b"AAAABBBBCCCCDDDD": %v`, err)
	}
	if !util.IsBytestringEqual(result, equal) {
		t.Fatalf(`TestGetPaddedBlock: getPaddedBlock(b"AAAABBBBCCCCDDDD" returns a modified block: %s`, util.BytestringAsString(result))
	}
}

func TestPaddingEmpty(t *testing.T) {
	lenZero := make([]byte, 0)
	expected := make([]byte, 16)
	result, err := getPaddedBlock(lenZero, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPaddingEmpty: Error when processing bytestring b"": %v`, err)
	}
	if !util.IsBytestringEqual(result, expected) {
		t.Fatalf(`TestPaddingEmpty: getPaddedBlock(b"") returns unexpected result "%s"`, util.BytestringAsString(result))
	}


	arbitrary := []byte("Hello world")
	expected = []byte{
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x77, 0x6F,
		0x72, 0x6C, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	result, err = getPaddedBlock(arbitrary, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPaddingEmpty: Error when processing bytestring b"Hello world": %v`, err)
	}
	if !util.IsBytestringEqual(result, expected) {
		t.Fatalf(`TestPaddingEmpty: getPaddedBlock(b"Hello world") returns unexpected result "%s"`, util.BytestringAsString(result))
	}
	
}
