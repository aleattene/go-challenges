package solution_dna_to_rna_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDNAtoRNA(t *testing.T) {
	assert.Equal(t, DNAtoRNA("GCATT"), "GCAUU")
	assert.Equal(t, DNAtoRNA("TGCAT"), "UGCAU")
	assert.Equal(t, DNAtoRNA("GTCAT"), "GUCAU")
	assert.Equal(t, DNAtoRNA("CCATCTTAGGGGGCGTGTAGGTGCA"), "CCAUCUUAGGGGGCGUGUAGGUGCA")
}
