// Package to "strain" lists based on a given function and either keep or discard the results.
package strain

// Ints is a slice of ints
type Ints []int

// Lists is a slice os slice of ints
type Lists [][]int

// Strings is a slice of strings
type Strings []string

// Keep is performed on a given Ints and returns a new Ints with only ints that satisfy the parameter function.
func (ints Ints) Keep(satisfies func(int) bool) Ints {
	var keeps Ints
	for _, val := range ints {
		if satisfies(val) {
			keeps = append(keeps, val)
		}
	}

	return keeps
}

// Discard is performed on a given Ints and returns a new Ints with only ints that do not satisfy the parameter function.
func (ints Ints) Discard(satisfies func(int) bool) Ints {
	var keeps Ints
	for _, val := range ints {
		if !satisfies(val) {
			keeps = append(keeps, val)
		}
	}

	return keeps
}

// Keep is performed on a given Lists and returns a new Lists with only []ints that satisfy the parameter function.
func (lists Lists) Keep(satisfies func([]int) bool) Lists {
	var keeps Lists
	for _, val := range lists {
		if satisfies(val) {
			keeps = append(keeps, val)
		}
	}

	return keeps
}

// Keep is performed on a given Strings and returns a new Strings with only strings that satisfy the parameter function.
func (strings Strings) Keep(satisfies func(string) bool) Strings {
	var keeps Strings
	for _, val := range strings {
		if satisfies(val) {
			keeps = append(keeps, val)
		}
	}

	return keeps
}
