package domain

import (
	"testing"

	"github.com/getmiranda/meli-challenge-api/domain/humans"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	dst := []interface{}{
		humans.Human{},
	}
	assert.EqualValues(t, dst, Migrate())
}
