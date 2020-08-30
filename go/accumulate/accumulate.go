package accumulate

// Accumulate change an element of a collection based on the required operation
func Accumulate(w []string, oper func(string) string) []string {
	converted := make([]string, len(w))
	for i, x := range w {
		converted[i] = oper(x)
	}
	return converted
}
