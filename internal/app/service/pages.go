package service

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/internal/repository"
)

func GetPagesByDiaryId(diaryId, userId uint) ([]models.Page, error) {
	return repository.GetPagesByDiaryId(diaryId, userId)
}

func GetPageById(pageId, userId uint) (models.Page, error) {
	return repository.GetPageById(pageId, userId)
}

func CreatePage(page *models.Page, userId uint) error {
	return repository.CreatePage(page, userId)
}

func UpdatePage(page *models.Page, userId uint) error {
	return repository.UpdatePage(page, userId)
}

func DeletePage(pageId, userId uint) error {
	return repository.DeletePage(pageId, userId)
}
