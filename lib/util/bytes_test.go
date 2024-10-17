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
