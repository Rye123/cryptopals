package aes

import (
	"testing"

	"github.com/Rye123/cryptopals/lib/util"
)

func TestPadUnpadBlock(t *testing.T) {
	// Test invalid length
	invalid := []byte("AAAABBBBCCCCDDDDe")
	_, err := padBlock(invalid, AESPadding_EMPTY)
	if err == nil {
		t.Fatalf("TestPadUnpadBlock (pad): Expected error from invalid bytestring of length %d", len(invalid))
	}
	_, err = unpadBlock(invalid, AESPadding_EMPTY)
	if err == nil {
		t.Fatalf("TestPadUnpadBlock (unpad): Expected error from invalid bytestring of length %d", len(invalid))
	}

	// Test equal length
	equal := []byte("AAAABBBBCCCCDDDD")
	result, err := padBlock(equal, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPadUnpadBlock: Error when padding bytestring b"AAAABBBBCCCCDDDD": %v`, err)
	}
	if !util.IsBytestringEqual(result, equal) {
		t.Fatalf(`TestPadUnpadBlock: padBlock(b"AAAABBBBCCCCDDDD" returns a modified block: %s`, util.BytestringAsString(result))
	}
	result, err = unpadBlock(equal, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPadUnpadBlock: Error when unpadding bytestring b"AAAABBBBCCCCDDDD": %v`, err)
	}
	if !util.IsBytestringEqual(result, equal) {
		t.Fatalf(`TestPadUnpadBlock: unpadBlock(b"AAAABBBBCCCCDDDD" returns a modified block: %s`, util.BytestringAsString(result))
	}
}

func TestPaddingEmpty(t *testing.T) {
	lenZero := make([]byte, 0)
	expected := make([]byte, 16)
	result, err := padBlock(lenZero, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPaddingEmpty: Error when padding bytestring b"": %v`, err)
	}
	if !util.IsBytestringEqual(result, expected) {
		t.Fatalf(`TestPaddingEmpty: padBlock(b"") returns unexpected result "%s"`, util.BytestringAsString(result))
	}
	result, err = unpadBlock(expected, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPaddingEmpty: Error when unpadding bytestring 0x0000000 00000000: %v`, err)
	}
	if !util.IsBytestringEqual(result, lenZero) {
		t.Fatalf(`TestPaddingEmpty: unpadBlock(0x00000000 00000000) returns unexpected result "%s"`, util.BytestringAsString(result))
	}


	arbitrary := []byte("Hello world")
	expected = []byte{
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x77, 0x6F,
		0x72, 0x6C, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	result, err = padBlock(arbitrary, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPaddingEmpty: Error when processing bytestring b"Hello world": %v`, err)
	}
	if !util.IsBytestringEqual(result, expected) {
		t.Fatalf(`TestPaddingEmpty: padBlock(b"Hello world") returns unexpected result "%s"`, util.BytestringAsString(result))
	}
	result, err = unpadBlock(expected, AESPadding_EMPTY)
	if err != nil {
		t.Fatalf(`TestPaddingEmpty: Error when unpadding bytestring b"Hello world\x00\x00\x00\x00\x00": %v`, err)
	}
	if !util.IsBytestringEqual(result, arbitrary) {
		t.Fatalf(`TestPaddingEmpty: unpadBlock(b"Hello world\x00\x00\x00\x00\x00") returns unexpected result "%s"`, util.BytestringAsString(result))
	}
	
}
