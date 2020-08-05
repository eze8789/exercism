package hamming

import (
	"errors"
)

// Distance measure the Hamming distance between two DNA string samples
func Distance(a, b string) (int, error) {

	aRune, bRune := []rune(a), []rune(b)

	// If not the same lenght exit
	if len(aRune) != len(bRune) {
		return 0, errors.New("The DNA samples have different lenght")
	}

	// Initialize variable to count distance
	var c = 0

	for i := range aRune {
		if aRune[i] != bRune[i] {
			c++
		}
	}
	return c, nil
}
