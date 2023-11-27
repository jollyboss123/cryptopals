package basics

import "fmt"

func fixedXOR(input string) ([]byte, error) {
	o, err := decodeHex(input)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("decoded: %s\n", o)
	fmt.Println(o)
	return o, nil
}
