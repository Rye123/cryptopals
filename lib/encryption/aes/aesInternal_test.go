package aes

import (
	"testing"
	"github.com/Rye123/cryptopals/lib/util"
)

func TestAESErrors(t *testing.T) {
	// Test invalid data length
	_, err := AESEncrypt([]byte("asdf"), []byte("AAAABBBBCCCCDDDD"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESEncrypt(invalid data length) gave no error")
	}
	_, err = AESDecrypt([]byte("asdf"), []byte("AAAABBBBCCCCDDDD"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESDecrypt(invalid data length) gave no error")
	}
	
	// Test invalid key lengths
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte(""))
	if err == nil {
		t.Fatalf("TestAESErrors: AESEncrypt(empty pass) gave no error")
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte(""))
	if err == nil {
		t.Fatalf("TestAESErrors: AESDecrypt(empty pass) gave no error")
	}
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESEncrypt(1-len key) gave no error")
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESDecrypt(1-len key) gave no error")
	}
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDE"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESEncrypt(15-len key) gave no error")
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDE"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESDecrypt(15-len key) gave no error")
	}
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEFG"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESEncrypt(17-len key) gave no error")
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEFG"))
	if err == nil {
		t.Fatalf("TestAESErrors: AESDecrypt(17-len key) gave no error")
	}

	// Test appropriate lengths
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEF"))
	if err != nil {
		t.Fatalf("TestAESErrors: AESEncrypt(16-len key) gave error %v", err)
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEF"))
	if err != nil {
		t.Fatalf("TestAESErrors: AESDecrypt(16-len key) gave error %v", err)
	}
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEFGHIJKLMO"))
	if err != nil {
		t.Fatalf("TestAESErrors: AESEncrypt(24-len key) gave error %v", err)
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEFGHIJKLMO"))
	if err != nil {
		t.Fatalf("TestAESErrors: AESDecrypt(24-len key) gave error %v", err)
	}
	_, err = AESEncrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEFGHIJKLMOPQRSTUVW"))
	if err != nil {
		t.Fatalf("TestAESErrors: AESEncrypt(32-len key) gave error %v", err)
	}
	_, err = AESDecrypt([]byte("AAAABBBBCCCCDDDD"), []byte("0123456789ABCDEFGHIJKLMOPQRSTUVW"))
	if err != nil {
		t.Fatalf("TestAESErrors: AESDecrypt(32-len key) gave error %v", err)
	}
}

func TestAESEncryptDecrypt(t *testing.T) {
	tests_bytestr := [][]byte{
		[]byte("0000000000000000"),
		[]byte("0000000000000000"),
		[]byte("0000000000000000"),
		[]byte("TESTING ONE TWO "),
		[]byte("The quick brown "),
		[]byte("lorem ipsum dolo"),
		[]byte("AAAABBBBCCCCDDDD"),
	}

	tests_key := [][]byte{
		[]byte("passwordpassword"),
		[]byte("passwordpasswordpassword"),
		[]byte("passwordpasswordpasswordpassword"),
		[]byte("passwordpassword"),
		{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe, 0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0xba, 0xbe},
		{0xab, 0xcd, 0xbe, 0xef, 0xab, 0xcd, 0xba, 0xbe, 0xba, 0xdc, 0xbe, 0xef, 0xba, 0xdc, 0xba, 0xbe},
		{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
	}

	expcs := [][]byte{
		{0x47, 0xa8, 0xe6, 0xf1, 0xb5, 0xa8, 0x1a, 0x37, 0x02, 0xd0, 0x30, 0x3a, 0x88, 0x47, 0x36, 0xc9},
		{0xa0, 0xbf, 0xdb, 0x4c, 0x5d, 0x71, 0x78, 0xe0, 0xef, 0xb7, 0x8b, 0x85, 0xb6, 0x20, 0x0e, 0x78},
		{0x6f, 0x53, 0x67, 0xfc, 0x68, 0x72, 0xdf, 0x67, 0x3d, 0xaf, 0xc2, 0xcf, 0xce, 0xe9, 0x13, 0xf0},
		{0xd9, 0x56, 0xeb, 0x62, 0x74, 0x13, 0xe3, 0xf3, 0xe4, 0x75, 0xe4, 0x38, 0xb1, 0x7d, 0xdf, 0x47},
		{0x53, 0x0c, 0xd9, 0x2f, 0xa7, 0x7b, 0x77, 0xe9, 0x84, 0x7d, 0x55, 0xc3, 0x9d, 0xec, 0xd4, 0xf3},
		{0x06, 0x40, 0xef, 0xe6, 0x0f, 0x34, 0x07, 0x89, 0xb1, 0x47, 0x3a, 0x37, 0xc1, 0xa8, 0xc5, 0xa2},
		{0x77, 0xd1, 0x13, 0x17, 0xab, 0x42, 0x0d, 0x27, 0xc9, 0xb4, 0x76, 0x42, 0x5d, 0x3c, 0x5d, 0x95},
	}

	if len(tests_bytestr) != len(tests_key) || len(tests_key) != len(expcs) {
		t.Fatalf(`TestAESEncryptDecrypt: Test error: tests count do not match`)
	}

	for i := 0; i < len(expcs); i++ {
		result, err := AESEncrypt(tests_bytestr[i], tests_key[i])
		if err != nil {
			t.Fatalf(`TestAESEncryptDecrypt: AESEncrypt(tests_bytestr[%d], tests_key[%d]) gave error: %v`, i, i, err)
		}
		if !util.IsBytestringEqual(result, expcs[i]) {
			t.Fatalf(`TestAESEncryptDecrypt: AESEncrypt(tests_bytestr[%d], tests_key[%d]) != expcs[%d]; expected: "%s"; got: "%s"`, i, i, i, util.BytestringAsString(expcs[i]), util.BytestringAsString(result))
		}

		decrypted, err := AESDecrypt(result, tests_key[i])
		if err != nil {
			t.Fatalf(`TestAESEncryptDecrypt: AESDecrypt(result[%d], tests_key[%d]) gave error: %v`, i, i, err)
		}
		if !util.IsBytestringEqual(tests_bytestr[i], decrypted) {
			t.Fatalf(`TestAesEncryptDecrypt: AESDecrypt(result[%d], tests_key[%d]) != tests_bytestr[%d]; expected: "%s"; got: "%s"`, i, i, i, util.BytestringAsString(tests_bytestr[i]), util.BytestringAsString(decrypted))
		}
	}
}
