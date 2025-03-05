package db

import (
	models2 "ArchiveOfBeing/internal/app/models"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models2.User{},
		&models2.Diary{},
		&models2.Character{},
		&models2.Page{},
	)

	if err != nil {
		return err
	}

	return nil
}
