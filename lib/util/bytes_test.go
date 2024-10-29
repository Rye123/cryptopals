package util

import "testing"

func TestHammingDistance(t *testing.T) {
	tests_h1 := [][]byte{
		[]byte(""),
		[]byte("abcd"),
		[]byte("a"),
		[]byte("the  "),
		[]byte("this is a test"),
		[]byte("the quick brown fox jumps over the lazy dog"),
	}

	tests_h2 := [][]byte{
		[]byte(""),
		[]byte("abcd"),
		[]byte(" "),
		[]byte("quick"),
		[]byte("wokka wokka!!!"),
		[]byte("lorem ipsum dolor sit amet, consectetur adi"),
	}

	expcs := []int{
		0,
		0,
		2,
		15,
		37,
		127,
	}

	if len(tests_h1) != len(tests_h2) || len(tests_h2) != len(expcs) {
		t.Fatalf(`TestHammingDistance: Test error: tests count do not match`)
	}

	for i := 0; i < len(expcs); i++ {
		distance, err := HammingDistance(tests_h1[i], tests_h2[i])
		if err != nil {
			t.Fatalf(`TestHammingDistance: test failed with unexpected error: %v`, err)
		}
		if distance != expcs[i] {
			t.Fatalf(`TestHammingDistance: test failed for HammingDistance(tests_h1[%d], tests_h2[%d]), expected %d, was %d`, i, i, expcs[i], distance)
		}
	}

	// error checks
	test_h1_err := []byte("test")
	test_h2_err := []byte("tes")
	distance, err := HammingDistance(test_h1_err, test_h2_err)
	if err == nil {
		t.Fatalf(`TestHammingDistance: expected length error from comparison of %s and %s, got no error with distance %d`, test_h1_err, test_h2_err, distance)
	}
}

func TestCircularRotate(t *testing.T) {
	// Ensure no modification of original data
	dat := []byte("test")
	result := CircularRotate(dat, 1, true)
	if !IsBytestringEqual(result, []byte("estt")) {
		t.Fatalf(`TestCircularRotate: CircularRotate(b"test", 1, true) != b"estt"`)
	}
	if !IsBytestringEqual(dat, []byte("test")) {
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
		if !IsBytestringEqual(result, expcs[i]) {
			t.Fatalf(`TestCircularRotate: results[%d] != expcs[%d]. results[%d] = "%s", expcs[%d] = "%s"`, i, i, i, result, i, expcs[i])
		}
	}
}
