package repository

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/pkg/db"
	"ArchiveOfBeing/pkg/errs"
	"ArchiveOfBeing/pkg/logger"
	"gorm.io/gorm"
)

func GetPagesByDiaryId(diaryId, userId uint) ([]models.Page, error) {
	_, err := GetDiaryByID(diaryId, userId)
	if err != nil {
		return nil, err
	}

	var pages []models.Page
	if err := db.GetDBConn().Where("diary_id = ?", diaryId).Find(&pages).Error; err != nil {
		logger.Error.Printf("[repository.GetPageByDiaryId] Error getting pages by diary id %d, user id %d: %v", diaryId, userId, err)
		return nil, TranslateGormError(err)
	}

	return pages, nil
}

func GetPageById(pageId, userId uint) (models.Page, error) {
	var page models.Page
	if err := db.GetDBConn().Where("id = ?", pageId).First(&page).Error; err != nil {
		logger.Error.Printf("[repository.GetPageById] Error getting page by id %v", pageId)
		return page, TranslateGormError(err)
	}

	var _ models.Diary
	_, err := GetDiaryByID(page.DiaryID, userId)
	if err != nil {
		return models.Page{}, TranslateGormError(err)
	}

	return page, nil
}

func CreatePage(page *models.Page, userId uint) error {
	tx := db.GetDBConn().Begin() // Начинаем транзакцию
	if tx.Error != nil {
		return tx.Error
	}

	character, err := GetCharacterByID(page.CharacterID, userId)
	if err != nil {
		tx.Rollback()
		return TranslateGormError(err)
	}

	if character.DiaryID != page.DiaryID {
		tx.Rollback()
		return errs.ErrCharacterNotFound
	}

	pages, err := GetPagesByDiaryId(page.DiaryID, userId)
	if err != nil {
		tx.Rollback()
		return TranslateGormError(err)
	}

	page.PageNumber = uint(len(pages) + 1)

	// Создаем страницу в транзакции
	if err := tx.Create(page).Error; err != nil {
		tx.Rollback() // Откат в случае ошибки
		logger.Error.Printf("[repository.CreatePage] Error creating page %v: %v", page, err)
		return TranslateGormError(err)
	}

	// Восстанавливаем листы дневника в рамках той же транзакции
	err = RestoreDiarySheets(tx, page.DiaryID, userId)
	if err != nil {
		tx.Rollback() // Откат в случае ошибки
		logger.Error.Printf("[repository.CreatePage] Error restoring diary sheets: %v", err)
		return TranslateGormError(err)
	}

	// Фиксируем изменения, если все прошло успешно
	if err := tx.Commit().Error; err != nil {
		logger.Error.Printf("[repository.CreatePage] Error committing transaction: %v", err)
		return err
	}

	return nil
}

func UpdatePage(page *models.Page, userId uint) error {
	tx := db.GetDBConn().Begin()
	diary, err := GetDiaryByIdUtil(tx, page.DiaryID, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if diary.UserID != userId {
		tx.Rollback()
		return errs.ErrDiaryNotFound
	}

	character, err := GetCharacterByID(page.CharacterID, userId)
	if err != nil {
		tx.Rollback()
		return TranslateGormError(err)
	}

	if character.DiaryID != page.DiaryID {
		tx.Rollback()
		return errs.ErrCharacterNotFound
	}

	pages, err := GetPagesByDiaryId(page.DiaryID, userId)
	if err != nil {
		tx.Rollback()
		return TranslateGormError(err)
	}

	page.PageNumber = uint(len(pages) + 1)

	if err := tx.Save(page).Error; err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.UpdatePage] Error updating page %v: %v", page, err)
		return TranslateGormError(err)
	}

	if tx.Commit().Error != nil {
		logger.Error.Printf("[repository.UpdatePage] Error committing transaction: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func DeletePage(pageId, userId uint) error {
	tx := db.GetDBConn().Begin()
	var page models.Page
	if err := tx.Where("id = ?", pageId).First(&page).Error; err != nil {
		tx.Rollback()
		return TranslateGormError(err)
	}

	diary, err := GetDiaryByIdUtil(tx, page.DiaryID, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if diary.UserID != userId {
		tx.Rollback()
		return errs.ErrDiaryNotFound
	}

	if err := tx.Delete(&page).Error; err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.DeletePage] Error deleting page %d: %v", pageId, err)
		return TranslateGormError(err)
	}

	if err := RestoreDiarySheets(tx, page.DiaryID, userId); err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.DeletePage] Error restoring diary sheets: %v", err)
		return TranslateGormError(err)
	}

	if err := restorePagesNums(tx, diary, pageId); err != nil {
		tx.Rollback()
		logger.Error.Printf("[repository.DeletePage] Error restoring pages nums: %v", err)
		return TranslateGormError(err)
	}

	if err := tx.Commit().Error; err != nil {
		logger.Error.Printf("[repository.DeletePage] Error committing transaction: %v", err)
		return TranslateGormError(err)
	}

	return nil
}

func restorePagesNums(tx *gorm.DB, diary models.Diary, deletedPageId uint) error {
	var pages []models.Page
	if err := tx.Where("diary_id = ? AND id != ?", diary.ID, deletedPageId).Find(&pages).Error; err != nil {
		return err
	}

	var cntPages uint
	for _, page := range pages {
		cntPages++
		page.PageNumber = cntPages
		if err := tx.Save(&page).Error; err != nil {
			tx.Rollback()
			return TranslateGormError(err)
		}
	}

	return nil
}
