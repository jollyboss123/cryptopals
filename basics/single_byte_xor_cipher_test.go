package basics

import (
	"reflect"
	"testing"
)

func TestFindXORCipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expect := []byte("Cooking MC's like a pound of bacon")
	res, err := findXORCipher(input)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(res, expect) {
		t.Fatal()
	}
}
