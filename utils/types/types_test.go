package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	assert.EqualValues(t, "X-Request-Id", XRequestId)
}

func TestMarshalFloat(t *testing.T) {
	f := Float(1.23)
	s, err := f.MarshalJSON()
	assert.Nil(t, err)
	assert.EqualValues(t, "1.23", string(s))

	f = Float(1.23456789)
	s, err = f.MarshalJSON()
	assert.Nil(t, err)
	assert.EqualValues(t, "1.23", string(s))

	f = Float(1)
	s, err = f.MarshalJSON()
	assert.Nil(t, err)
	assert.EqualValues(t, "1.00", string(s))

	f = Float(0)
	s, err = f.MarshalJSON()
	assert.Nil(t, err)
	assert.EqualValues(t, "0.00", string(s))
}
