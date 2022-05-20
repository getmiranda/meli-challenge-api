package errors_utils

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessageError(t *testing.T) {
	result := getMessageError("error")
	assert.EqualValues(t, "error", result)

	result = getMessageError(errors.New("error"))
	assert.EqualValues(t, "error", result)

	result = getMessageError(nil)
	assert.EqualValues(t, "<nil>", result)
}

func TestRestErr(t *testing.T) {
	err := &restErr{
		ErrStatus: http.StatusBadRequest,
		ErrError:  "error",
	}

	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeError(t *testing.T) {
	err := MakeError(http.StatusBadRequest, "error")
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeBadRequestError(t *testing.T) {
	err := MakeBadRequestError("error")
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeTooEarlyError(t *testing.T) {
	err := MakeTooEarlyError("error")
	assert.EqualValues(t, http.StatusTooEarly, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeNotFoundError(t *testing.T) {
	err := MakeNotFoundError("error")
	assert.EqualValues(t, http.StatusNotFound, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeBadGatewayError(t *testing.T) {
	err := MakeBadGatewayError("error")
	assert.EqualValues(t, http.StatusBadGateway, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeInternalServerError(t *testing.T) {
	err := MakeInternalServerError("error")
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "error", err.Error())
}

func TestMakeErrorFromBytes(t *testing.T) {
	apiErr, err := MakeErrorFromBytes([]byte(`{"status":400,"error":"error"}`))
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "error", apiErr.Error())

	apiErr, err = MakeErrorFromBytes([]byte(``))
	assert.Nil(t, apiErr)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid json", err.Error())
}
