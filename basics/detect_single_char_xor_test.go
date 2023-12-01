package basics

import (
	"reflect"
	"testing"
)

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
