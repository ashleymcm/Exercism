package strand

// ToRNA takes a DNA strand, represented by a string, and converts it to RNA
func ToRNA(dna string) string {
	var rna string

	for i := range dna {
		rna += Complement(dna[i])
	}

	return rna
}

// Complement takes a nucleotide, represented by a string, and returns its complement
func Complement(nucleotide byte) string {
	var complement = map[byte]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}

	return complement[nucleotide]
}
