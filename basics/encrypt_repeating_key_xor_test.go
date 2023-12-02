package basics

import (
	"reflect"
	"testing"
)

// TestEncryptRepeatingKeyXOR validates https://cryptopals.com/sets/1/challenges/5.
func TestEncryptRepeatingKeyXOR(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	expect := []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

	res, err := encryptRepeatingKeyXOR(input, key)
	if err != nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(expect, res) {
		t.Fatal()
	}
}
