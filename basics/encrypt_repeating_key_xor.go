package basics

import "fmt"

// encryptRepeatingKeyXOR resolves the set 1 challenge 5: implement repeating-key XOR
// at https://cryptopals.com/sets/1/challenges/5. Encrypts the input string by using
// repeating-key XOR.
func encryptRepeatingKeyXOR(input, key string) ([]byte, error) {
	src := []byte(input)
	ek := createRepeatingKey(key, len(src))
	res := xor(src, ek)
	fmt.Printf("%s\n", res)
	fmt.Printf("%s\n", encodeHexBytes(res))
	return encodeHexBytes(res), nil
}

// createRepeatingKey creates and returns a repeating key of a given size.
func createRepeatingKey(key string, size int) []byte {
	dst := make([]byte, size)
	for i := 0; i < size; i++ {
		dst[i] = key[i%len(key)]
	}
	return dst
}
