// Package hamming to calculate Hamming distance
package hamming

import "errors"

// Distance calculates Hamming distance between two strands of DNA.
// Returns 0 and error if the two strands cannot be compared.
func Distance(a, b string) (int, error) {
	// If strings are different lengths they cannot be compared
	if len(a) != len(b) {
		return 0, errors.New("cannot compare strands of differing lengths")
	}

	var length = len(a)
	var hamming int
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			hamming++
		}
	}

	return hamming, nil
}
