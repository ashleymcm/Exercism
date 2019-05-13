// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb controls the printing of a proverb.
package proverb

import "fmt"

// Proverb takes an array of strings and returns an array of sentences that
// form a proverb.
func Proverb(rhyme []string) []string {
	var proverb = []string{}
	var length = len(rhyme)

	if length > 0 {
		for i := 0; i < length-1; i++ {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}

		proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return proverb
}
