package basics

import (
	"fmt"
)

// fixedXOR resolves the set 1 solution of fixed xor
// at https://cryptopals.com/sets/1/challenges/2. Takes two hexadecimal string input,
// decodes it into bytes, and xor against each other, then encodes it into a hexadecimal string.
func fixedXOR(input, xorInput string) ([]byte, error) {
	o, err := decodeHex(input)
	if err != nil {
		return nil, err
	}
	j, err := decodeHex(xorInput)
	if err != nil {
		return nil, err
	}
	res := xor(o, j)
	fmt.Println(res)
	return encodeHexBytes(res), nil
}

// xor runs the xor bitwise operator on the two input byte arrays
func xor(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}

// encodeHexBytes takes a byte slice and converts it to a hexadecimal string.
// It iterates over each byte in the byte slice and convert into two
// hexadecimal characters.
func encodeHexBytes(src []byte) []byte {
	dst := make([]byte, len(src)*2)
	for i, b := range src {
		hexChars := byteToHexChar(b)
		dst[i*2] = hexChars[0]
		dst[i*2+1] = hexChars[1]
	}
	return dst
}

// byteToHexChar converts a single byte to its hexadecimal character equivalent
func byteToHexChar(c byte) []byte {
	const hexChars = "0123456789abcdef"
	return []byte{hexChars[c>>4], hexChars[c&0x0f]}
}
