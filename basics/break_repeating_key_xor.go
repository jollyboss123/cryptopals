package basics

import (
	"errors"
)

// hamming returns the edit distance between a and b. The hamming
// distance is defined by the number of bits that are different between two strings.
func hamming(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("input strings must have the same length")
	}
	dist := 0

	// xor the two values and count the number of 1-bits in the result
	x := xor([]byte(a), []byte(b))
	for i := range x {
		diff := x[i]

		for j := 0; j < 8; j++ {
			dist += int(diff & 1)
			diff >>= 1
		}
	}
	return dist, nil
}
