package controllers

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/internal/app/service"
	"ArchiveOfBeing/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCharacterByDiaryID(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	diaryId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	characters, err := service.GetCharacterByDiaryAndUserID(diaryId, userId)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": characters})
}

func GetCharacterByID(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	characterId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	character, err := service.GetCharacterByID(characterId, userId)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": character})
}

func CreateCharacter(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	var character models.Character
	if err := c.ShouldBindJSON(&character); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	character.UserID = userId

	if err := service.CreateCharacter(&character); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Character created successfully!"})
}

func UpdateCharacter(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	characterId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	var character models.Character
	if err := c.ShouldBindJSON(&character); err != nil {
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	character.UserID = userId
	character.ID = characterId

	if err := service.UpdateCharacter(&character); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "character updated successfully!"})
}

func DeleteCharacter(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		HandleError(c, err)
		return
	}

	characterId, err := getParam(c, "id")
	if err != nil {
		HandleError(c, err)
		return
	}

	if err := service.DeleteCharacter(characterId, userId); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Character deleted successfully!"})
}
