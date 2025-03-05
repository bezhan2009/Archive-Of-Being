package service

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/internal/repository"
	"ArchiveOfBeing/pkg/errs"
	"errors"
)

func GetAllUserDiaries(userId uint) ([]models.Diary, error) {
	return repository.GetAllUserDiaries(userId)
}

func GetDiaryByID(diaryId uint, userID uint) (models.Diary, error) {
	return repository.GetDiaryByID(diaryId, userID)
}

func CreateDiary(diary *models.Diary, userID uint) error {
	diary.UserID = userID
	if _, err := repository.GetDiaryByTitle(diary.Title, userID); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return repository.CreateDiary(diary)
		}
		return err
	} else {
		return errs.ErrUniquenessFailed
	}
}

func UpdateDiary(diary *models.Diary, userID uint) error {
	diary.UserID = userID
	if _, err := repository.GetDiaryByTitle(diary.Title, userID); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return repository.UpdateDiary(diary)
		}
		return err
	} else {
		return errs.ErrUniquenessFailed
	}
}

func DeleteDiary(diaryId uint, userID uint) error {
	return repository.DeleteDiary(diaryId, userID)
}
