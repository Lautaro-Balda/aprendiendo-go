package strand

func ToRNA(dna string) string {
	var rna []rune
	for _, val := range dna {
		switch val {
		case 'A':
			rna = append(rna, 'U')
		case 'G':
			rna = append(rna, 'C')
		case 'C':
			rna = append(rna, 'G')
		case 'T':
			rna = append(rna, 'A')
		}
	}
	return string(rna)
}
