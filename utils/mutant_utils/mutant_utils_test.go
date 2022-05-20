package mutant_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSquare(t *testing.T) {
	result := IsSquare([]string{"a"})
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
		"OOOO",
		"OOOO",
		"OOOO",
		"OOOO",
	})
	assert.EqualValues(t, false, result.IsMutant)

	result = IsMutant([]string{
		"OOOOOO",
		"OOOOOO",
		"OOOOOA",
		"OOOOOA",
		"OOOOOA",
		"OOOOOA",
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
		"OOOOOO",
		"OOAOOO",
		"OOOAOO",
		"OOOOAO",
		"OOOOOA",
		"OOOOOO",
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
		"OOOOOOOOA",
		"OOOOOOOAO",
		"OOOOOOAOO",
		"OOOOOAOOO",
		"OOOOAOOOO",
		"OAOAOAAAO",
		"OOAAAAOOO",
		"OAOAOOOOO",
		"OOOOAOAOO",
	})

	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)

	result = IsMutant([]string{
		"OOOOOO",
		"OOOOOO",
		"OOOOOA",
		"OOOOAO",
		"OOOAOO",
		"OOAOOO",
	})

	assert.EqualValues(t, true, result.IsMutant)
	assert.EqualValues(t, "AAAA", result.Dna)
}
