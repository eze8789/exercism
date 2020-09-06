package isogram

import (
	"unicode"
)

// IsIsogram check if a giving word is an isogram. Ref.: https://en.wikipedia.org/wiki/Isogram
func IsIsogram(w string) bool {
	letters := make(map[rune]int)
	for _, v := range w {
		_, ok :=  letters[unicode.ToLower(v)]
		if unicode.IsLetter(v) {
			letters[unicode.ToLower(v)] = 1
		}
		if ok {
			return false
		}
	}
	return true
}
