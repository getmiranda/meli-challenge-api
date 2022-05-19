package errors_utils

import (
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
	Causes() []interface{}
}

type restErr struct {
	ErrStatus int           `json:"status"`
	ErrError  string        `json:"error"`
	ErrCauses []interface{} `json:"-"`
}

func (v *restErr) Status() int {
	return v.ErrStatus
}

func (v *restErr) Error() string {
	return v.ErrError
}

func (e *restErr) Causes() []interface{} {
	return e.ErrCauses
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
func MakeInternalServerError(v interface{}, err ...error) RestErr {
	result := &restErr{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  getMessageError(v),
	}

	if len(err) > 0 {
		result.ErrCauses = make([]interface{}, len(err))
		for i, e := range err {
			result.ErrCauses[i] = e
		}
	}

	return result
}
