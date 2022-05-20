package matrix_utils

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
