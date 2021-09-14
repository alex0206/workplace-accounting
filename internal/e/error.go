package e

import (
	"fmt"
	"net/http"
)

// Error describe error methods
type Error interface {
	Code() int
	Detail() string
	Error() string
}

type httpError struct {
	detail string
	code   int
}

func (e httpError) Error() string {
	return fmt.Sprintf(`code: %d, detail: '%s'`, e.code, e.detail)
}

func (e httpError) Code() int {
	return e.code
}

func (e httpError) Detail() string {
	return e.detail
}

// NewInternal getting internal server error
func NewInternal(detail string) Error {
	return httpError{
		detail: detail,
		code:   http.StatusInternalServerError,
	}
}

// NewInternalf formatted internal server error
func NewInternalf(template string, args ...interface{}) Error {
	return httpError{
		detail: fmt.Sprintf(template, args...),
		code:   http.StatusInternalServerError,
	}
}

// NewBadRequest getting bad request error
func NewBadRequest(detail string) Error {
	return httpError{
		detail: detail,
		code:   http.StatusBadRequest,
	}
}
