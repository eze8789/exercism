package scrabble

import (
	"unicode"
)

// Score calculate how much point a word sum in scrabble game
func Score(word string) int {
	var c int
	for _, r := range word {
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			c++
		case 'D', 'G':
			c += 2
		case 'B', 'C', 'M', 'P':
			c += 3
		case 'F', 'H', 'V', 'W', 'Y':
			c += 4
		case 'K':
			c += 5
		case 'J', 'X':
			c += 8
		case 'Q', 'Z':
			c += 10
		}
	}
	return c
}
