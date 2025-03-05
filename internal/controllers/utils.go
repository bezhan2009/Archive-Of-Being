package controllers

import (
	"ArchiveOfBeing/internal/controllers/middlewares"
	"ArchiveOfBeing/pkg/errs"
	"github.com/gin-gonic/gin"
	"strconv"
)

func getParam(c *gin.Context, param string) (uint, error) {
	gotParamStr := c.Param(param)
	if gotParamStr == "" {
		return 0, errs.MissingParam(param)
	}

	gotParam, err := strconv.Atoi(gotParamStr)
	if err != nil {
		return 0, errs.InvalidParam(param)
	}

	return uint(gotParam), nil
}

func getUserID(c *gin.Context) (uint, error) {
	// Получаем ID пользователя
	userID := c.GetUint(middlewares.UserIDCtx)
	if userID == 0 {
		return 0, errs.ErrUnauthorized
	}

	return userID, nil
}
