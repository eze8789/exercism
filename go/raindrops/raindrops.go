package raindrops

import (
	"strconv"
)

// Convert change a number to some words depending if 3, 5 or 7 is a factor of it.
func Convert(n int) string {
	var r string

	if n%3 == 0 {
		r += "Pling"
	}
	if n%5 == 0 {
		r += "Plang"
	}
	if n%7 == 0 {
		r += "Plong"
	}
	if r == "" {
		r = strconv.Itoa(n)
	}
	return r
}
