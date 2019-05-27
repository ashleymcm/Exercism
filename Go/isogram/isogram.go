// Package isogram tests if a given string is an isogram
package isogram

import "unicode"

// IsIsogram tests if the given word is an isogram
func IsIsogram(word string) bool {
	var tally = map[rune]struct{}{}

	for _, letter := range word {
		letter = unicode.ToLower(letter)

		if letter == '-' || letter == ' ' {
			continue
		} else if _, exists := tally[letter]; exists {
			return false
		}

		tally[letter] = struct{}{}
	}

	return true
}
