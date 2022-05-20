package errors_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const ()

var (
	ErrDatabase = errors.New("database error")
)

type RestErr interface {
	Status() int
	Error() string
}

type restErr struct {
	ErrStatus int    `json:"status"`
	ErrError  string `json:"error"`
}

func (v *restErr) Status() int {
	return v.ErrStatus
}

func (v *restErr) Error() string {
	return v.ErrError
}

func getMessageError(v interface{}) string {
	if msg, ok := v.(string); ok {
		return msg
	} else if err, ok := v.(error); ok {
		return err.Error()
	} else {
		return fmt.Sprintf("%v", v)
	}
}

// MakeError returns a new error.
//
// type of v is string or error.
func MakeError(statusCode int, v interface{}) RestErr {
	return &restErr{
		ErrStatus: statusCode,
		ErrError:  getMessageError(v),
	}
}

// MakeBadRequestError returns a new bad request error.
//
// type of v is string or error.
func MakeBadRequestError(v interface{}) RestErr {
	return &restErr{
		ErrStatus: http.StatusBadRequest,
		ErrError:  getMessageError(v),
	}
}

// MakeTooEarlyError returns a new too early error.
//
// type of v is string or error.
func MakeTooEarlyError(v interface{}) RestErr {
	return &restErr{
		ErrStatus: http.StatusTooEarly,
		ErrError:  getMessageError(v),
	}
}

// MakeNotFoundError returns a new not found error.
//
// type of v is string or error.
func MakeNotFoundError(v interface{}) RestErr {
	return &restErr{
		ErrStatus: http.StatusNotFound,
		ErrError:  getMessageError(v),
	}
}

// MakeBadGatewayError returns a new bad gateway error.
//
// type of v is string or error.
func MakeBadGatewayError(v interface{}) RestErr {
	return &restErr{
		ErrStatus: http.StatusBadGateway,
		ErrError:  getMessageError(v),
	}
}

// MakeInternalServerError returns a new internal server error.
//
// type of v is string or error.
func MakeInternalServerError(v interface{}) RestErr {
	result := &restErr{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  getMessageError(v),
	}

	return result
}

// MakeErrorFromBytes returns a new error from bytes.
func MakeErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return &apiErr, nil
}
