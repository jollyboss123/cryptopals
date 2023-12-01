package basics

import (
	"errors"
	"fmt"
)

/**
A hexadecimal string represents arbitrary binary data with each digit having
one of 16 possible values. This equates to exactly 4 bits of data per digit.
Consequently, a pair of hex digits can represent a single 8-bit byte.

A Base64 string is another format for representing arbitrary binary data.
Each digit in Base64 can have one of 64 values, encoding 6 bits of data. This
results in a more compact representation than hexadecimal.
*/

// convertHexToBase64 resolves the set 1 challenge 1: convert hex to base64
// at https://cryptopals.com/sets/1/challenges/1. Takes a hexadecimal string input,
// decodes it into bytes, and encodes it into a base64 string.
func convertHexToBase64(input string) ([]byte, error) {
	o, err := decodeHex(input)
	if err != nil {
		return nil, err
	}
	fmt.Printf("decoded: %s\n", o)
	return encodeToBase64(o), nil
}

// decodeHex takes a hexadecimal string and converts it to a byte slice.
// It iterates over pairs of characters in the src string and combines
// these two 4-bit values into one 8-bit byte.
func decodeHex(src string) ([]byte, error) {
	if len(src)%2 != 0 {
		return nil, errors.New("invalid length: hex string must have an even number of characters")
	}

	dst := make([]byte, len(src)/2)

	for i := 0; i < len(src); i += 2 {
		a, err := hexCharToByte(src[i])
		if err != nil {
			return nil, err
		}
		b, err := hexCharToByte(src[i+1])
		if err != nil {
			return nil, err
		}
		dst[i/2] = (a << 4) | b
	}

	return dst, nil
}

// hexCharToByte converts a single hexadecimal character to its byte equivalent.
func hexCharToByte(c byte) (byte, error) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', nil
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, nil
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, nil
	}
	return 0, fmt.Errorf("invalid character: %c", c)
}

const (
	// base64Chars defines the encoding table for base64.
	base64Chars      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	pad         rune = '-'
)

// encodeToBase64 converts a byte slice into a base64 encoded string.
func encodeToBase64(src []byte) []byte {
	dst := make([]byte, (len(src)+2)/3*4) // each 3 byte block is encoded into 4 characters.

	i, j := 0, 0
	for ; i < len(src); i += 3 {
		// combine three 8-bit bytes into a 24-bit number and encode it
		val := uint(src[i])<<16 | uint(src[i+1])<<8 | uint(src[i+2])

		// encode the 24-bit number into four 6-bit base64 characters
		dst[j] = base64Chars[val>>18&0x3F]
		dst[j+1] = base64Chars[val>>12&0x3F]
		dst[j+2] = base64Chars[val>>6&0x3F]
		dst[j+3] = base64Chars[val&0x3F]

		j += 4
	}

	// handle any remaining bytes and add necessary padding
	if remain := len(src) - i; remain > 0 {
		val := uint(src[i]) << 16
		if remain == 2 {
			val |= uint(src[i+1]) << 8
		}

		dst[j] = base64Chars[val>>18&0x3F]
		dst[j+1] = base64Chars[val>>12&0x3F]

		switch remain {
		case 2:
			dst[j+2] = base64Chars[val>>6&0x3F]
			dst[j+3] = byte(pad)
		case 1:
			dst[j+2] = byte(pad)
			dst[j+3] = byte(pad)
		}
	}

	return dst
}
