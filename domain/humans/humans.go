package humans

import (
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
