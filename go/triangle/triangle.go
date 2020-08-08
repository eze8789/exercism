package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides calculate what type of triangle is it, based on it's sides.
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	var k Kind

	sort.Float64s(sides)

	switch {
	case math.IsNaN(sides[0]) || sides[0] <= 0 || sides[2] == math.Inf(1) ||
		sides[1]+sides[0] < sides[2]:
		k = NaT
	case sides[0] == sides[2]:
		k = Equ
	case sides[0] == sides[1] || sides[1] == sides[2]:
		k = Iso
	default:
		k = Sca
	}
	return k
}
