package basics

import (
	"reflect"
	"testing"
)

// TestDetectSingleCharXOR validates https://cryptopals.com/sets/1/challenges/4
func TestDetectSingleCharXOR(t *testing.T) {
	inputFile := "4.txt"
	expect := []byte("Now that the party is jumping\n")
	res, err := detectSingleCharXOR(inputFile)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(expect, res) {
		t.Fatal()
	}
}
