package aes

import (
	"testing"
	"github.com/Rye123/cryptopals/lib/util"
)

func TestCircularRotate(t *testing.T) {
	// Ensure no modification of original data
	dat := []byte("test")
	result := CircularRotate(dat, 1, true)
	if !util.IsBytestringEqual(result, []byte("estt")) {
		t.Fatalf(`TestCircularRotate: CircularRotate(b"test", 1, true) != b"estt"`)
	}
	if !util.IsBytestringEqual(dat, []byte("test")) {
		t.Fatalf(`TestCircularRotate: CircularRotate modifies original data`)
	}
	
	// Main test
	tests := [][]byte{
		{},
		[]byte("test"),
		[]byte("test"),
		{0xca, 0xfe, 0xba},
		{0xca, 0xfe, 0xba, 0xbe, 0xde, 0xad},
		[]byte("test"),
		{0xaa, 0xbb, 0xcc, 0xdd},
		{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe},
	}
	offsets := []int{
		2,
		0,
		3,
		3,
		4,
		1,
		2,
		6,
	}
	dirLefts := []bool{
		true,
		true,
		true,
		true,
		true,
		false,
		false,
		false,
	}
	expcs := [][]byte{
		{},
		[]byte("test"),
		[]byte("ttes"),
		{0xca, 0xfe, 0xba},
		{0xde, 0xad, 0xca, 0xfe, 0xba, 0xbe},
		[]byte("ttes"),
		{0xcc, 0xdd, 0xaa, 0xbb},
		{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe},
	}

	if len(tests) != len(offsets) || len(tests) != len(dirLefts) || len(tests) != len(expcs) {
		t.Fatalf("TestCircularRotate: Invalid length of lists")
	}
	for i := 0; i < len(tests); i++ {
		result := CircularRotate(tests[i], offsets[i], dirLefts[i])
		if !util.IsBytestringEqual(result, expcs[i]) {
			t.Fatalf(`TestCircularRotate: results[%d] != expcs[%d]. results[%d] = "%s", expcs[%d] = "%s"`, i, i, i, result, i, expcs[i])
		}
	}
}
