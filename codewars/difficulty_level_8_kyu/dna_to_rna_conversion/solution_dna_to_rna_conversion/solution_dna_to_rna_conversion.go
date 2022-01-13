package solution_dna_to_rna_conversion

import "strings"

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}
