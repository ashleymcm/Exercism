// Package to control the way a string about sharing prints.
package twofer
	
import "fmt"

// Prints who is being shared with. If string is empty, default is "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
