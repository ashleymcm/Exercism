package accumulate

// Accumulate takes an array of strings and a function (that accepts and returns a string)
// and returns an array of strings having that function performed on each
func Accumulate(strings []string, action func(string) string) []string {
	for i := 0; i < len(strings); i++ {
		strings[i] = action(strings[i])
	}

	return strings
}
