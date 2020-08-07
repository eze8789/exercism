package scrabble

import (
	"strings"
	"unicode"
)

const onePoints = "A, E, I, O, U, L, N, R, S, T"
const twoPoints = "D, G"
const threePoints = "B, C, M, P"
const fourPoints = "F, H, V, W, Y"
const fivePoints = "K"
const eightPoints = "J, X"
const tenPoints = "Q, Z"

// Score calculate how much point a word sum in scrabble game
func Score(word string) int {
	var c int
	for _, r := range word {
		r = unicode.ToUpper(r)
		switch {
		case strings.ContainsRune(onePoints, r):
			c++
		case strings.ContainsRune(twoPoints, r):
			c += 2
		case strings.ContainsRune(threePoints, r):
			c += 3
		case strings.ContainsRune(fourPoints, r):
			c += 4
		case strings.ContainsRune(fivePoints, r):
			c += 5
		case strings.ContainsRune(eightPoints, r):
			c += 8
		case strings.ContainsRune(tenPoints, r):
			c += 10
		}
	}
	return c
}
