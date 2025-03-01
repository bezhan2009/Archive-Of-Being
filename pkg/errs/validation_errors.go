package errs

import "errors"

// Validation Errors
var (
	ErrValidationFailed    = errors.New("ErrValidationFailed")
	ErrInvalidToken        = errors.New("ErrInvalidToken")
	ErrRefreshTokenExpired = errors.New("ErrRefreshTokenExpired")
)
