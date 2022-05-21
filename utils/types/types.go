package types

import "fmt"

type Key string

const (
	XRequestId Key = "X-Request-Id"
)

type Float float64

func (f Float) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%.2f", f)
	return []byte(s), nil
}
