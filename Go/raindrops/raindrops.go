package raindrops

import "fmt"

// Convert converts an integer into a raindrop-string.
func Convert(num int) string {
	var drop = ""

	if num%3 == 0 {
		drop += "Pling"
	}
	if num%5 == 0 {
		drop += "Plang"
	}
	if num%7 == 0 {
		drop += "Plong"
	}

	if drop == "" {
		drop += fmt.Sprint(num)
	}

	return drop
}
