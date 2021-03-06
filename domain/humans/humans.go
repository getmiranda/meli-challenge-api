package humans

import (
	"strings"

	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/getmiranda/meli-challenge-api/utils/mutant_utils"
	"github.com/getmiranda/meli-challenge-api/utils/types"
	"gorm.io/gorm"
)

type Human struct {
	gorm.Model
	Dna      string `gorm:"column:dna"`
	IsMutant bool   `gorm:"column:is_mutant;index"`
}

// TableName overrides the table name used by Human to `humans`.
func (Human) TableName() string {
	return "humans"
}

type HumanRequest struct {
	Dna []string `json:"dna"`
}

// Validate validates the HumanRequest.
func (s *HumanRequest) Validate() errors_utils.RestErr {
	if len(s.Dna) == 0 {
		return errors_utils.MakeBadRequestError("dna is required")
	}

	if !mutant_utils.IsSquare(s.Dna) {
		return errors_utils.MakeBadRequestError("dna must be a square matrix")
	}

	if !mutant_utils.IsValidDna(s.Dna) {
		return errors_utils.MakeBadRequestError("dna must be composed only of 'A', 'T', 'C' and 'G'")
	}

	return nil
}

// GenerateDna generates the dna as a string from the HumanRequest.
func (s *HumanRequest) GenerateDna() string {
	return strings.Join(s.Dna, "-")
}

type StatsResponse struct {
	CountMutantDna int64       `json:"count_mutant_dna"`
	CountHumanDna  int64       `json:"count_human_dna"`
	Ratio          types.Float `json:"ratio"`
}
