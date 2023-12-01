package basics

import (
	"fmt"
)

// fixedXOR resolves the set 1 challenge 2: fixed xor
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

// xor runs the xor bitwise operator on the two input byte arrays.
// If both bytes input are of different length, will take the max length between
// the two inputs, but assume that trailing 0 bytes will be truncated.
func xor(a, b []byte) []byte {
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	c := make([]byte, n)
	for i := 0; i < len(a) && i < len(b); i++ {
		c[i] = a[i] ^ b[i]
	}
	for i := len(b); i < len(a); i++ {
		c[i] = a[i]
	}
	for i := len(a); i < len(b); i++ {
		c[i] = b[i]
	}
	s := len(c)
	for range c {
		if s > 0 && c[s-1] == 0 {
			s--
		}
	}
	if s < len(c) {
		dst := make([]byte, s)
		copy(dst, c)
		return dst
	}
	return c
}

// encodeHexBytes takes a byte slice and converts it to a hexadecimal string.
// It iterates over each byte in the byte slice and convert into two
// hexadecimal characters.
func encodeHexBytes(src []byte) []byte {
	dst := make([]byte, len(src)*2)
	for i, b := range src {
		hc := byteToHexChar(b)
		dst[i*2] = hc[0]
		dst[i*2+1] = hc[1]
	}
	return dst
}

// hexChars defines the encoding table for hexadecimal.
const hexChars = "0123456789abcdef"

// byteToHexChar converts a single byte to its hexadecimal character equivalent.
func byteToHexChar(c byte) []byte {
	return []byte{hexChars[c>>4], hexChars[c&0x0f]}
}
