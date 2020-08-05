// Package twofer is for share with another person
package twofer

// ShareWith is the function to share with another person.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
