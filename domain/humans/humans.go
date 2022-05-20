package humans

import (
	"github.com/getmiranda/meli-challenge-api/utils/errors_utils"
	"github.com/getmiranda/meli-challenge-api/utils/mutant_utils"
	"gorm.io/gorm"
)

type Human struct {
	gorm.Model
	Dna      string `gorm:"column:dna"`
	IsMutant bool   `gorm:"column:is_mutant"`
}

// TableName overrides the table name used by Human to `humans`.
func (Human) TableName() string {
	return "humans"
}

type HumanRequest struct {
	Dna []string `json:"dna"`
}

func (s *HumanRequest) Validate() errors_utils.RestErr {
	if len(s.Dna) == 0 {
		return errors_utils.MakeBadRequestError("dna is required")
	}

	if !mutant_utils.IsSquare(s.Dna) {
		return errors_utils.MakeBadRequestError("dna must be a square matrix")
	}

	return nil
}
