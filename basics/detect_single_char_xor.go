package basics

import (
	"fmt"
	"os"
	"strings"
)

// detectSingleCharXOR resolves the set 1 challenge 4: detect single-character XOR
// at https://cryptopals.com/sets/1/challenges/3. Reads the input file and
// run findXORCipher on each line of the input file while keeping track of the best score.
func detectSingleCharXOR(filename string) ([]byte, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lns := strings.Split(string(f), "\n")
	var res []byte
	var score float64
	for _, l := range lns {
		d, s, err := findXORCipher(l)
		if err != nil {
			return nil, err
		}
		if s > score {
			res = d
			score = s
		}
	}
	fmt.Printf("decrypted: %s", res)
	fmt.Printf("score: %f\n", score)
	return res, nil
}
