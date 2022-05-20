package humans

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	assert.EqualValues(t, "humans", Human{}.TableName())
}

func TestValidate(t *testing.T) {
	t.Run("ErrorDnaIsEmpty", func(t *testing.T) {
		input := HumanRequest{}
		err := input.Validate()

		assert.NotNil(t, err)
		assert.EqualValues(t, "dna is required", err.Error())
		assert.EqualValues(t, http.StatusBadRequest, err.Status())
	})

	t.Run("ErrorDnaIsNotSquare", func(t *testing.T) {
		input := HumanRequest{
			Dna: []string{
				"ATGCGA",
				"CAGTGC",
				"TTATGT",
				"AGAAGG",
				"CCCCTA",
				"TCACTGW",
			},
		}
		err := input.Validate()

		assert.NotNil(t, err)
		assert.EqualValues(t, "dna must be a square matrix", err.Error())
		assert.EqualValues(t, http.StatusBadRequest, err.Status())
	})

	t.Run("Success", func(t *testing.T) {
		input := HumanRequest{
			Dna: []string{
				"ATGCGA",
				"CAGTGC",
				"TTATTT",
				"AGACGG",
				"GCGTCA",
				"TCACTG",
			},
		}
		err := input.Validate()

		assert.Nil(t, err)
	})
}

func TestGenerateDna(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		input := HumanRequest{
			Dna: []string{
				"ATGCGA",
				"CAGTGC",
				"TTATTT",
				"AGACGG",
				"GCGTCA",
				"TCACTG",
			},
		}
		dna := input.GenerateDna()

		assert.EqualValues(t, "ATGCGA-CAGTGC-TTATTT-AGACGG-GCGTCA-TCACTG", dna)
	})
}
