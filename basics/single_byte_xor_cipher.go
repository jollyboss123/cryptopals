package basics

import (
	"fmt"
	"math"
)

// findXORCipher resolves the set 1 challenge 3: single-byte XOR cipher
// at https://cryptopals.com/sets/1/challenges/3. Takes an input and finds the
// single-byte XOR cipher determined by the scoring of the XOR'ed result based on
// character frequency.
func findXORCipher(input string) ([]byte, float64, error) {
	o, err := decodeHex(input)
	if err != nil {
		return nil, 0, err
	}
	fmt.Printf("decoded: %s\n", o)
	var res []byte
	var score float64
	for i := 0; i < 256; i++ {
		r := make([]byte, len(o))
		for j := range r {
			r[j] = byte(i)
		}
		x := xor(o, r)
		s := englishness(x)
		if s > score {
			res = x
			score = s
		}
		s = 0
	}
	fmt.Printf("decrypted: %s\n", res)
	fmt.Printf("score: %f\n", score)
	return res, score, nil
}

// englishness determines how likely a byte array is to be english based on
// character frequency. Returns a value between 0 and 1, where 0 means 'totally
// unlike english' and 1 means 'totally like english'. Uses the Bhattacharyya coefficient.
func englishness(src []byte) float64 {
	frequencytable := map[byte]float64{
		byte('a'): 0.08167,
		byte('b'): 0.01492,
		byte('c'): 0.02782,
		byte('d'): 0.04253,
		byte('e'): 0.1270,
		byte('f'): 0.0228,
		byte('g'): 0.02015,
		byte('h'): 0.06094,
		byte('i'): 0.06966,
		byte('j'): 0.00153,
		byte('k'): 0.00772,
		byte('l'): 0.04025,
		byte('m'): 0.02406,
		byte('n'): 0.06749,
		byte('o'): 0.07507,
		byte('p'): 0.01929,
		byte('q'): 0.00095,
		byte('r'): 0.05987,
		byte('s'): 0.06327,
		byte('t'): 0.09056,
		byte('u'): 0.02758,
		byte('v'): 0.00978,
		byte('w'): 0.02360,
		byte('x'): 0.00150,
		byte('y'): 0.01974,
		byte('z'): 0.00074,
	}

	var coefficient float64
	for char, count := range counter(string(src)) {
		freq, ok := frequencytable[byte(char)]
		if ok {
			coefficient += math.Sqrt(freq * (float64(count) / float64(len(src))))
		}
	}

	return coefficient
}

// counter calculates each character frequency of a given input
// and returns a map of each character as key and its frequency as value.
func counter(input string) map[rune]int {
	freq := make(map[rune]int)
	for _, v := range input {
		freq[v] = freq[v] + 1
	}

	return freq
}
