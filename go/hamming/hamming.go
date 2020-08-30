package hamming

import (
	"errors"
)

// Distance measure the Hamming distance between two DNA string samples
func Distance(a, b string) (int, error) {

	aRune, bRune := []rune(a), []rune(b)

	if len(aRune) != len(bRune) {
		return 0, errors.New("the dna samples have different length")
	}

	var c = 0

	for i := range aRune {
		if aRune[i] != bRune[i] {
			c++
		}
	}
	return c, nil
}
