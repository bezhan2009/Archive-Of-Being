package errs

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func MissingParam(paramName string) error {
	return errors.New(fmt.Sprintf("Missing %s", paramName))
}

func InvalidParam(paramName string) error {
	return errors.New(fmt.Sprintf("Invalid param %s", paramName))
}

func UniquenessError(modelObj string, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRecordNotFound
	}

	return errors.New(fmt.Sprintf("%s with this name is already exists", modelObj))
}

// Validation Errors
var (
	ErrValidationFailed    = errors.New("ErrValidationFailed")
	ErrInvalidToken        = errors.New("ErrInvalidToken")
	ErrRefreshTokenExpired = errors.New("ErrRefreshTokenExpired")
	ErrUniquenessFailed    = errors.New("ErrUniquenessFailed")
)
