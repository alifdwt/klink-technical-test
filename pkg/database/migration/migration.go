package migration

import (
	"technical-test/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Member{},
		&models.Gender{},
		&models.Level{},
	)

	if err != nil {
		panic(err)
	}

	return err
}
