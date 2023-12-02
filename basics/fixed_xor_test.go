package basics

import (
	"reflect"
	"testing"
)

// TestFixedXOR validates https://cryptopals.com/sets/1/challenges/2
func TestFixedXOR(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	expect := []byte("746865206b696420646f6e277420706c6179")
	xorAgainst := "686974207468652062756c6c277320657965"
	res, err := fixedXOR(input, xorAgainst)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(expect, res) {
		t.Fatal()
	}
}
