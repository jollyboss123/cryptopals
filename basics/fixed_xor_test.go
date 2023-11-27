package basics

import (
	"reflect"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	expect := "746865206b696420646f6e277420706c6179"
	res, err := fixedXOR(input)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(expect, res) {
		t.Fatal()
	}
}
