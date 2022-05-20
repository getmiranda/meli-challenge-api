package mutant_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSquare(t *testing.T) {
	result := IsSquare([]string{})
	assert.EqualValues(t, false, result)

	result = IsSquare([]string{"a"})
	assert.EqualValues(t, true, result)

	result = IsSquare([]string{"a", "b"})
	assert.EqualValues(t, false, result)

	result = IsSquare([]string{"aaaa"})
	assert.EqualValues(t, false, result)
}

func TestIsMutant(t *testing.T) {
	result := IsMutant([]string{
		"ATGCTA",
		"CTGTGC",
		"TTATGT",
		"AGAAGG",
		"CCTCTA",
		"TCACTG",
	})
	assert.EqualValues(t, false, result.IsMutant)

	result = IsMutant([]string{
		"ATGCTA",
		"CTGTGC",
		"TTATGT",
		"AGAAGG",
		"CCTCTA",
		"TCTTTT",
	})
	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "TTTT", result.Dna)
	assert.EqualValues(t, []Coordinate{
		{5, 2},
		{5, 3},
		{5, 4},
		{5, 5},
	}, result.Coordinates)

	result = IsMutant([]string{
		"GAGG",
		"ATOA",
		"GCGG",
		"GAGG",
	})
	assert.EqualValues(t, false, result.IsMutant)

	result = IsMutant([]string{
		"TAGTCG",
		"AGGTGG",
		"CGTGGA",
		"GTACGA",
		"GGCGTA",
		"GAGCGA",
	})
	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)
	assert.EqualValues(t, []Coordinate{
		{2, 5},
		{3, 5},
		{4, 5},
		{5, 5},
	}, result.Coordinates)

	result = IsMutant([]string{
		"TCTGCG",
		"GGAGTT",
		"AGGAGG",
		"GCGGAG",
		"GGGCTA",
		"TGCGGG",
	})
	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)
	assert.EqualValues(t, []Coordinate{
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}, result.Coordinates)

	result = IsMutant([]string{
		"AGGATGAGA",
		"GTCGTTGAG",
		"GCTGCGATG",
		"TGCTTAGCC",
		"TGCGACGTG",
		"CAGAGAAAG",
		"TGAAAAGGC",
		"GAGAGGCAG",
		"GCGGAGAGG",
	})

	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)

	result = IsMutant([]string{
		"AGTC",
		"AGTC",
		"AGTC",
		"AGTC",
	})

	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)

	result = IsMutant([]string{
		"AGTA",
		"CGAC",
		"GATC",
		"AGTC",
	})

	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)
}

func TestIsValidDna(t *testing.T) {
	result := IsValidDna([]string{
		"ATGCTA",
		"CTGTGC",
		"TTATGT",
		"AGAAGG",
		"CCTCTA",
		"TCACTG",
	})
	assert.EqualValues(t, true, result)

	result = IsValidDna([]string{
		"ATGCTA",
		"CTGTGC",
		"TTATGT",
		"AGAAGG",
		"CCTCTA",
		"TCACTQ",
	})
	assert.EqualValues(t, false, result)
}
