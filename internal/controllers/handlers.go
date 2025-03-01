package controllers

import (
	"ArchiveOfBeing/pkg/errs"
	"ArchiveOfBeing/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Обработка ошибок, которые приводят к статусу 400 (Bad Request)
func handleBadRequestErrors(err error) bool {
	return errors.Is(err, errs.ErrUsernameUniquenessFailed) ||
		errors.Is(err, errs.ErrIncorrectUsernameOrPassword) ||
		errors.Is(err, errs.ErrInvalidField) ||
		errors.Is(err, errs.ErrEmailIsEmpty) ||
		errors.Is(err, errs.ErrPasswordIsEmpty) ||
		errors.Is(err, errs.ErrUsernameIsEmpty) ||
		errors.Is(err, errs.ErrDeleteFailed)
}

// Обработка ошибок, которые приводят к статусу 404 (Not Found)
func handleNotFoundErrors(err error) bool {
	return errors.Is(err, errs.ErrRecordNotFound)
}

// Обработка ошибок, которые приводят к статусу 401 (Unauthorized)
func handleUnauthorizedErrors(err error) bool {
	return errors.Is(err, errs.ErrInvalidToken) ||
		errors.Is(err, errs.ErrUnauthorized) ||
		errors.Is(err, errs.ErrRefreshTokenExpired)
}

// HandleError Основная функция обработки ошибок
func HandleError(c *gin.Context, err error) {
	if handleBadRequestErrors(err) {
		c.JSON(http.StatusBadRequest, newErrorResponse(err.Error()))
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
