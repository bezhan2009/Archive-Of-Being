package repository

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/pkg/db"
	"ArchiveOfBeing/pkg/errs"
	"ArchiveOfBeing/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func GetAllUserDiaries(userId uint) ([]models.Diary, error) {
	var userDiaries []models.Diary
	if err := db.GetDBConn().Where("user_id = ?", userId).Find(&userDiaries).Error; err != nil {
		logger.Error.Printf("[repository.GetAllUserDiaries] Error getting user diaries: %s", err.Error())
		return nil, TranslateGormError(err)
	}

	return userDiaries, nil
}

func GetDiaryByIdUtil(tx *gorm.DB, diaryId, userId uint) (models.Diary, error) {
	var diary models.Diary
	if err := tx.Where("id = ? AND user_id = ?", diaryId, userId).First(&diary).Error; err != nil {
		logger.Error.Printf("[repository.GetDiaryByIdUtil] Error getting diary by ID: %s", err.Error())
		return diary, TranslateGormError(err)
	}

	return diary, nil
}

func GetDiaryByID(diaryId, userId uint) (models.Diary, error) {
	var diary models.Diary
	if err := db.GetDBConn().Where("id = ? AND user_id = ?", diaryId, userId).First(&diary).Error; err != nil {
		logger.Error.Printf("[repository.GetDiaryByIdUtil] Error getting diary by ID: %s", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return diary, errs.ErrDiaryNotFound
		}

		return diary, TranslateGormError(err)
	}

	return diary, nil
}

func GetDiaryByTitle(diaryTitle string, userId uint) (models.Diary, error) {
	var diary models.Diary
	if err := db.GetDBConn().Where("title = ? AND user_id = ?", diaryTitle, userId).First(&diary).Error; err != nil {
		logger.Error.Printf("[repository.GetDiaryByTitle] Error getting diary by Title: %s", err.Error())
		return diary, TranslateGormError(err)
	}

	return diary, nil
}

func CreateDiary(diary *models.Diary) error {
	if err := db.GetDBConn().Create(diary).Error; err != nil {
		logger.Error.Printf("[repository.CreateDiary] Error creating diary: %s", err.Error())
		return TranslateGormError(err)
	}

	return nil
}

func UpdateDiary(diary *models.Diary) error {
	if err := db.GetDBConn().Save(diary).Error; err != nil {
		logger.Error.Printf("[repository.UpdateDiary] Error updating diary: %s", err.Error())
		return TranslateGormError(err)
	}

	return nil
}

func RestoreDiarySheets(tx *gorm.DB, diaryId, userId uint) error {
	diary, err := GetDiaryByIdUtil(tx, diaryId, userId) // Теперь используем транзакцию
	if err != nil {
		return TranslateGormError(err)
	}

	var pages []models.Page
	if err := tx.Where("diary_id = ?", diary.ID).Find(&pages).Error; err != nil {
		logger.Error.Printf("[repository.RestoreDiarySheets] Error getting pages: %s", err.Error())
		return TranslateGormError(err)
	}

	diary.Pages = uint(len(pages))

	if err := tx.Save(&diary).Error; err != nil {
		logger.Error.Printf("[repository.RestoreDiarySheets] Error updating diary sheets: %s", err.Error())
		return TranslateGormError(err)
	}

	return nil
}

func DeleteDiary(diaryId uint, userId uint) error {
	diary, err := GetDiaryByIdUtil(db.GetDBConn(), diaryId, userId)
	if err != nil {
		return TranslateGormError(err)
	}

	if err := db.GetDBConn().Delete(&diary).Error; err != nil {
		logger.Error.Printf("[repository.DeleteDiary] Error deleting diary: %s", err.Error())
		return TranslateGormError(err)
	}

	return nil
}
