package domain

import (
	"github.com/getmiranda/meli-challenge-api/domain/humans"
)

// AutoMigrate is a function that returns a list of models that needs to be migrated to the database.
func Migrate() []interface{} {
	dst := []interface{}{
		humans.Human{},
	}
	return dst
}
