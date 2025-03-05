package repository

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/pkg/db"
	"ArchiveOfBeing/pkg/logger"
)

func GetPageByDiaryId(diaryId, userId uint) ([]models.Page, error) {
	var pages []models.Page
	if err := db.GetDBConn().Where("diary_id = ? and user_id = ?", diaryId, userId).Find(&pages).Error; err != nil {
		logger.Error.Printf("[repository.GetPageByDiaryId] Error getting pages by diary id %d, user id %d: %v", diaryId, userId, err)
		return nil, TranslateGormError(err)
	}

	return pages, nil
}

func GetPageById(pageId, userId uint) (models.Page, error) {
	var page models.Page
	if err := db.GetDBConn().Where("id = ? and user_id = ?", pageId, userId).First(&page).Error; err != nil {
		logger.Error.Println("[repository.GetPageById] Error getting page by id %d", pageId)
		return page, TranslateGormError(err)
	}

	return page, nil
}

func CreatePage(page *models.Page) error {
	if err := db.GetDBConn().Create(page).Error; err != nil {
		logger.Error.Printf("[repository.CreatePage] Error creating page %d: %v", page, err)
		return TranslateGormError(err)
	}

	return nil
}

func UpdatePage(page *models.Page) error {
	if err := db.GetDBConn().Save(page).Error; err != nil {
		logger.Error.Printf("[repository.UpdatePage] Error updating page %d: %v", page, err)
		return TranslateGormError(err)
	}

	return nil
}

func DeletePage(pageId uint) error {
	if err := db.GetDBConn().Delete(&models.Page{}, "id = ?", pageId).Error; err != nil {
		logger.Error.Printf("[repository.DeletePage] Error deleting page %d: %v", pageId, err)
		return TranslateGormError(err)
	}

	return nil
}
