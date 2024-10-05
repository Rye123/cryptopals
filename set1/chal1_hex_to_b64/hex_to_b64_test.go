package hex_to_b64

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	tests := []string{
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"c18163beedae3119fcd5c40d3245d049",
		"b8e1d96423f143c6992544ba7ef2b07a16f4950705b0bacd0af2a130c3d2bc2c",
		"27381212586d80f80bfbcc7472d5ffe7ab09e5cb8784db91e759c76f2fd852bfc358358c720ff0e8d40d2f1481b5e194",
		"02c7de520a38",
		"",
	}
	expcs := []string{
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		"wYFjvu2uMRn81cQNMkXQSQ==",
		"uOHZZCPxQ8aZJUS6fvKwehb0lQcFsLrNCvKhMMPSvCw=",
		"JzgSElhtgPgL+8x0ctX/56sJ5cuHhNuR51nHby/YUr/DWDWMcg/w6NQNLxSBteGU",
		"AsfeUgo4",
		"",
	}
	for i, test := range tests {
		result, err := HexToBase64(test)
		if err != nil {
			t.Fatalf(`HexToBase64("%s") gave an error: %v`, test, err)
		}
		expected := expcs[i]
		if result != expected {
			t.Fatalf(`HexToBase64("%s") = "%s", expected "%s"`, test, result, expected)
		}
	}
}
