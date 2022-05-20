package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	assert.EqualValues(t, "X-Request-Id", XRequestId)
}
