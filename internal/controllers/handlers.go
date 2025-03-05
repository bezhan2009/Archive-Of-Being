package controllers

import (
	"ArchiveOfBeing/pkg/errs"
	"ArchiveOfBeing/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Обработка ошибок, приводящих к статусу 400 (Bad Request)
func handleBadRequestErrors(err error) bool {
	return errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrValidationFailed) ||
		errors.Is(err, errs.ErrInvalidField) ||
		errors.Is(err, errs.ErrEmailIsEmpty) ||
		errors.Is(err, errs.ErrUniquenessFailed) ||
		errors.Is(err, errs.ErrPasswordIsEmpty) ||
		errors.Is(err, errs.ErrUsernameIsEmpty) ||
		errors.Is(err, errs.ErrDeleteFailed) ||
		isCustomBadRequestError(err)
}

// Проверка ошибок, созданных функциями MissingParam, InvalidParam
func isCustomBadRequestError(err error) bool {
	errMsg := err.Error()
	return strings.HasPrefix(errMsg, "Missing ") || strings.HasPrefix(errMsg, "Invalid param ")
}

// Проверка ошибки уникальности
func isUniquenessError(err error) bool {
	return strings.Contains(err.Error(), "with this name is already exists")
}

// Обработка ошибок, приводящих к статусу 404 (Not Found)
func handleNotFoundErrors(err error) bool {
	return errors.Is(err, errs.ErrRecordNotFound) ||
		errors.Is(err, errs.ErrDiaryNotFound)
}

// Обработка ошибок, приводящих к статусу 401 (Unauthorized)
func handleUnauthorizedErrors(err error) bool {
	return errors.Is(err, errs.ErrInvalidToken) ||
		errors.Is(err, errs.ErrUnauthorized) ||
		errors.Is(err, errs.ErrRefreshTokenExpired)
}

// HandleError Основная функция обработки ошибок
func HandleError(c *gin.Context, err error) {
	if handleBadRequestErrors(err) {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
	} else if isUniquenessError(err) {
		c.JSON(http.StatusConflict, newErrorResponse(err.Error()))
	} else if errors.Is(err, errs.ErrPermissionDenied) {
		c.JSON(http.StatusForbidden, newErrorResponse(err.Error()))
	} else if handleNotFoundErrors(err) {
		c.JSON(http.StatusNotFound, newErrorResponse(err.Error()))
	} else if handleUnauthorizedErrors(err) {
		c.JSON(http.StatusUnauthorized, newErrorResponse(err.Error()))
	} else {
		logger.Error.Printf("Err: %s", err)
		c.JSON(http.StatusInternalServerError, newErrorResponse(errs.ErrSomethingWentWrong.Error()))
	}
}
