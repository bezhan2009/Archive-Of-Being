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
	)

	if err != nil {
		return err
	}

	return nil
}
