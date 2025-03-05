package service

import (
	"ArchiveOfBeing/internal/app/models"
	"ArchiveOfBeing/internal/repository"
	"ArchiveOfBeing/pkg/errs"
	"errors"
)

func GetCharacterByDiaryAndUserID(diaryID, userId uint) ([]models.Character, error) {
	return repository.GetCharacterByDiaryAndUserID(diaryID, userId)
}

func GetCharacterByID(characterID, userId uint) (models.Character, error) {
	return repository.GetCharacterByID(characterID, userId)
}

func CreateCharacter(character *models.Character) error {
	if _, err := repository.GetDiaryByID(character.DiaryID, character.UserID); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return errs.ErrDiaryNotFound
		}
	}

	if _, err := repository.GetCharacterByDiaryUserIdAndTitle(character.DiaryID, character.UserID, character.Title); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			return repository.CreateCharacter(character)
		}
		return err
	} else {
		return errs.ErrUniquenessFailed
	}
}

func UpdateCharacter(character *models.Character) error {
	if _, err := repository.GetCharacterByDiaryUserIdAndTitle(character.DiaryID, character.UserID, character.Title); err != nil {
		if errors.Is(err, errs.ErrRecordNotFound) {
			if _, err := repository.GetCharacterByUserIdAndCharacterId(character.UserID, character.ID); err != nil {
				return err
			} else {
				return repository.UpdateCharacter(character)
			}
		}
		return err
	} else {
		return errs.ErrUniquenessFailed
	}
}

func DeleteCharacter(characterId, userId uint) error {
	return repository.DeleteCharacter(characterId, userId)
}
