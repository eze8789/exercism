package etl

import (
	"strings"
)

// Transform change from an old scoring map (points to letters) to a new one (letters to points)
func Transform(i map[int][]string) map[string]int {
	o := make(map[string]int)
	for k, v := range i {
		for _, l := range v {
			o[strings.ToLower(l)]=k
		}
	}
	return o
}