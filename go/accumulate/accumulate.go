package accumulate

// Accumulate change an element of a collection based on the required operation
func Accumulate(w []string, oper func(string) string) []string {
	var converted []string
	for _, x := range w {
		converted = append(converted, oper(x))
	}
	return converted
}
