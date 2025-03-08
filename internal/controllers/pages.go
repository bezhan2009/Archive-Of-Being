package controllers

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/internal/app/service"
	"ArchiveOfBeing/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPagesByDiaryId(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Получаем ID пользователя
	diaryId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	pages, err := service.GetPagesByDiaryId(diaryId, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pages})
}

func GetPageById(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Получаем ID пользователя
	pageId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	page, err := service.GetPageById(userID, pageId)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": page})
}

func CreatePage(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	var page models.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	err = service.CreatePage(&page, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": page})
}

func UpdatePage(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Получаем ID пользователя
	pageId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	var page models.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	page.ID = pageId

	err = service.UpdatePage(&page, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Page updated successfully"})
}

func DeletePage(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	// Получаем ID пользователя
	pageId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	err = service.DeletePage(pageId, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Page deleted successfully"})
}
