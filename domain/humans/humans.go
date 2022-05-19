package humans

import (
	"gorm.io/gorm"
)

type Human struct {
	gorm.Model
	Dna      string `gorm:"column:dna"`
	IsMutant bool   `gorm:"column:is_mutant"`
}
