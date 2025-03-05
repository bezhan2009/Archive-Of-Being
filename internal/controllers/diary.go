package controllers

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/internal/app/service"
	"ArchiveOfBeing/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUserDiaries(c *gin.Context) {
	// Получаем ID пользователя
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	diaries, err := service.GetAllUserDiaries(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	if len(diaries) == 0 {
		c.JSON(http.StatusOK, gin.H{"data": "You have no diaries yet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diaries})
}

func GetDiaryByID(c *gin.Context) {
	// Получаем ID пользователя
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

	diary, err := service.GetDiaryByID(diaryId, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"diary": diary})
}

func CreateDiary(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	var diary models.Diary
	if err := c.ShouldBind(&diary); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	diary.UserID = userID

	if err := service.CreateDiary(&diary, userID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Diary created successfully"})
}

func UpdateDiary(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	var diary models.Diary
	if err := c.ShouldBind(&diary); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	// Получаем ID пользователя
	diaryId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	diary.ID = diaryId

	if err := service.UpdateDiary(&diary, userID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Diary updated successfully"})
}

func DeleteDiary(c *gin.Context) {
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

	if err = service.DeleteDiary(diaryId, userID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Diary deleted successfully"})
}
