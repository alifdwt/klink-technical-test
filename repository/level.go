package repository

import (
	"technical-test/models"

	"gorm.io/gorm"
)

type levelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) *levelRepository {
	return &levelRepository{db}
}

func (r *levelRepository) CreateLevel(name string) (*models.Level, error) {
	var level models.Level

	db := r.db.Model(&level)

	level.Name = name

	result := db.Debug().Create(&level)
	if result.Error != nil {
		return nil, result.Error
	}

	return &level, nil
}

func (r *levelRepository) GetLevelByName(name string) (*models.Level, error) {
	var level models.Level

	db := r.db.Model(&level)

	result := db.Debug().Where("name = ?", name).First(&level)
	if result.Error != nil {
		return nil, result.Error
	}

	return &level, nil
}

func (r *levelRepository) GetLevelById(id int) (*models.Level, error) {
	var level models.Level

	db := r.db.Model(&level)

	result := db.Debug().Where("id = ?", id).First(&level)
	if result.Error != nil {
		return nil, result.Error
	}

	return &level, nil
}
