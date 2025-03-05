package repository

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/pkg/db"
	"ArchiveOfBeing/pkg/errs"
	"ArchiveOfBeing/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func GetCharacterByDiaryAndUserID(diaryId, userId uint) ([]models.Character, error) {
	var characters []models.Character
	if err := db.GetDBConn().Where("diary_id = ? AND user_id = ?", diaryId, userId).Find(&characters).Error; err != nil {
		logger.Error.Printf("[repository.GetCharacterByDiaryAndUserID] Error getting character by diary_id: %d, user_id: %d: %v", diaryId, userId, err)
		return nil, TranslateGormError(err)
	}

	return characters, nil
}

func GetCharacterByID(characterId, userId uint) (models.Character, error) {
	var character models.Character // Изменил на единственное число
	result := db.GetDBConn().Where("id = ? AND user_id = ?", characterId, userId).First(&character)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			logger.Warn.Printf("[repository.GetCharacterByID] Character not found with id: %d, user_id: %d", characterId, userId)
			return models.Character{}, errs.ErrRecordNotFound // Возвращаем свою ошибку "не найдено"
		}
		logger.Error.Printf("[repository.GetCharacterByID] Error getting character by id: %d, user_id: %d: %v", characterId, userId, result.Error)
		return models.Character{}, TranslateGormError(result.Error)
	}

	return character, nil
}

func GetCharacterByDiaryUserIdAndTitle(diaryId, userId uint, title string) (models.Character, error) {
	var characters models.Character
	if err := db.GetDBConn().Where("diary_id = ? and title = ? AND user_id = ?", diaryId, title, userId).First(&characters).Error; err != nil {
		logger.Error.Printf("[repository.GetCharacterByDiaryUserIdAndTitle] Error getting character by diary_id: %d, user_id: %d: %v", diaryId, userId, err)
		return models.Character{}, TranslateGormError(err)
	}

	return characters, nil
}

func GetCharacterByUserIdAndCharacterId(userId, characterId uint) (models.Character, error) {
	var characters models.Character
	if err := db.GetDBConn().Where("user_id = ? AND id = ?", userId, characterId).First(&characters).Error; err != nil {
		logger.Error.Printf("[repository.GetCharacterByUserIdAndCharacterId] Error getting character by user_id: %d, character_id: %d: %v", userId, characterId, err)
		return models.Character{}, TranslateGormError(err)
	}

	return characters, nil
}

func CreateCharacter(character *models.Character) error {
	if err := db.GetDBConn().Create(&character).Error; err != nil {
		logger.Error.Printf("[repository.CreateCharacter] Error creating character: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func UpdateCharacter(character *models.Character) error {
	if err := db.GetDBConn().Model(&models.Character{}).Where("id = ?", character.ID).Update("title", character.Title).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCharacter] Error updating character: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeleteCharacter(characterId, userId uint) error {
	character, err := GetCharacterByUserIdAndCharacterId(userId, characterId)
	if err != nil {
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Delete(&character).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCharacter] Error deleting character: %v", err)
		return TranslateGormError(err)
	}

	return nil
}
