package set_1

import (
	"reflect"
	"testing"
)

// TestHexToBase64 validates https://cryptopals.com/sets/1/challenges/1
func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expect := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res, err := resolve(input)
	if err != nil {
		t.Fatal()
	}
	if reflect.DeepEqual(expect, res) {
		t.Fatal()
	}
}
