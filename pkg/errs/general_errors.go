package errs

import "errors"

// General Errors
var (
	ErrDiaryNotFound      = errors.New("ErrDiaryNotFound")
	ErrCharacterNotFound  = errors.New("ErrCharacterNotFound")
	ErrPageNotFound       = errors.New("ErrPageNotFound")
	ErrRecordNotFound     = errors.New("ErrRecordNotFound")
	ErrSomethingWentWrong = errors.New("ErrSomethingWentWrong")
	ErrUserNotFound       = errors.New("ErrUserNotFound")
	ErrDeleteFailed       = errors.New("ErrDeleteFailed")
	ErrInvalidData        = errors.New("ErrInvalidData")
)
