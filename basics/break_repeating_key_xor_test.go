package basics

import "testing"

func TestDecryptRepeatingKeyXOR(t *testing.T) {

}

func TestHamming(t *testing.T) {
	a := "this is a test"
	b := "wokka wokka!!!"
	expect := 37

	res, err := hamming(a, b)
	if err != nil {
		t.Fatal()
	}
	if res != expect {
		t.Fatal()
	}
}
