package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dmitrymomot/go-pkg/response"
)

// Predefined errors
var (
	ErrValidationFailed = errors.New("validation_failed")
	ErrNotFound         = errors.New("not_found")
	ErrSendCommand      = errors.New("send_command")
	ErrBindRequest      = errors.New("bind_request")
	ErrUnauthorized     = errors.New("unauthorized")
)

// Map error to http status code.
var errorStatus = map[error]int{
	ErrValidationFailed: http.StatusPreconditionFailed,
	ErrNotFound:         http.StatusNotFound,
	ErrSendCommand:      http.StatusInternalServerError,
	ErrBindRequest:      http.StatusBadRequest,
}

// Map error to error description.
var errorDescription = map[error]string{
	ErrValidationFailed: "Validation failed",
	ErrNotFound:         "Not found",
	ErrSendCommand:      "Something went wrong...",
	ErrBindRequest:      "Failed to bind request",
}

// ValidationError is a validation error.
type ValidationError struct {
	Code    int
	Err     error
	Message string
	Errors  url.Values
}

// Error returns error message.
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %v", e.Err, e.Errors)
}

// NewValidationError creates new validation error.
func NewValidationError(validationErr url.Values) ValidationError {
	return ValidationError{
		Code:    http.StatusPreconditionFailed,
		Err:     ErrValidationFailed,
		Message: "Validation failed",
		Errors:  validationErr,
	}
}

// NewError creates new error.
func NewError(err error) *response.Error {
	code := errorStatus[err]
	message := errorDescription[err]
	errx := http.StatusText(code)

	if code > 0 {
		return &response.Error{
			Code:    code,
			Error:   errx,
			Message: message,
		}
	}

	var validation url.Values

	if verr, ok := err.(ValidationError); ok {
		code = verr.Code
		errx = verr.Err.Error()
		message = verr.Message
		validation = verr.Errors
	}

	for kerr, statusCode := range errorStatus {
		if errors.Is(err, kerr) {
			code = statusCode
			errx = kerr.Error()
			message = err.Error()
		}
	}

	return &response.Error{
		Code:       code,
		Error:      errx,
		Message:    message,
		Validation: validation,
	}
}
