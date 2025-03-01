package errs

import "errors"

// General Errors
var (
	ErrRecordNotFound     = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong = errors.New("ErrSomethingWentWrong")
	ErrUserNotFound       = errors.New("ErrUserNotFound")
	ErrDeleteFailed       = errors.New("ErrDeleteFailed")
	ErrInvalidData        = errors.New("ErrInvalidData")
)
