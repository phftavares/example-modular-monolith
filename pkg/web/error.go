package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest           = Define("bad_request", http.StatusBadRequest)
	ErrUnauthorized         = Define("unauthorized", http.StatusUnauthorized)
	ErrForbidden            = Define("forbidden", http.StatusForbidden)
	ErrNotFound             = Define("not_found", http.StatusNotFound)
	ErrConflict             = Define("conflict", http.StatusConflict)
	ErrUnprocessableEntity  = Define("unprocessable_entity", http.StatusUnprocessableEntity)
	ErrUnsupportedMediaType = Define("unsupported_media_type", http.StatusUnsupportedMediaType)

	ErrInternalServerError = NewError(Define("internal_server_error", http.StatusInternalServerError), "Internal error. Please contact system admin!")
	ErrNotImplemented      = NewError(Define("not_implemented", http.StatusNotImplemented), "Operation not implemented!")
)

type Error struct {
	Code       string
	Message    string
	Details    interface{}
	statusCode int
}

func (err *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Details interface{} `json:"details,omitempty"`
	}{
		Code:    err.Code,
		Message: err.Message,
		Details: err.Details,
	})
}

func (err *Error) Error() string {
	return err.Message
}

func (err *Error) StatusCode() int {
	return err.statusCode
}

func Define(code string, statusCode int) *Error {
	return &Error{
		statusCode: statusCode,
		Code:       code,
	}
}

func NewError(err *Error, message string) *Error {
	err.Message = message
	return err
}

func NewErrorf(err *Error, message string, args ...interface{}) *Error {
	err.Message = fmt.Sprintf(message, args...)
	return err
}

func NewErrorWithDetails(err *Error, message string, details interface{}) *Error {
	err.Message = message
	err.Details = details
	return err
}
