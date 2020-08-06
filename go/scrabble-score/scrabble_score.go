package scrabble

import "strings"

const onePoints = "A, E, I, O, U, L, N, R, S, T"
const twoPoints = "D, G"
const threePoints = "B, C, M, P"
const fourPoints = "F, H, V, W, Y"
const fivePoints = "K"
const eightPoints = "J, X"
const tenPoints = "Q, Z"

// Score calculate how much point a word sum in scrabble game
func Score(word string) int {
	w := []rune(word)
	var c int

	for i := range w {
		l := strings.ToUpper(string(w[i]))
		switch {
		case strings.Contains(onePoints, l):
			c++
		case strings.Contains(twoPoints, l):
			c += 2
		case strings.Contains(threePoints, l):
			c += 3
		case strings.Contains(fourPoints, l):
			c += 4
		case strings.Contains(fivePoints, l):
			c += 5
		case strings.Contains(eightPoints, l):
			c += 8
		case strings.Contains(tenPoints, l):
			c += 10

		}
	}
	return c
}
