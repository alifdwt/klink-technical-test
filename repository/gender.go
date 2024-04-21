package repository

import (
	"technical-test/models"

	"gorm.io/gorm"
)

type genderRepository struct {
	db *gorm.DB
}

func NewGenderRepository(db *gorm.DB) *genderRepository {
	return &genderRepository{db}
}

func (r *genderRepository) CreateGender(name string) (*models.Gender, error) {
	var gender models.Gender

	db := r.db.Model(&gender)

	gender.Name = name

	result := db.Debug().Create(&gender)
	if result.Error != nil {
		return nil, result.Error
	}

	return &gender, nil
}

func (r *genderRepository) GetGenderById(id int) (*models.Gender, error) {
	var gender models.Gender

	db := r.db.Model(&gender)

	result := db.Debug().Where("id = ?", id).First(&gender)
	if result.Error != nil {
		return nil, result.Error
	}

	return &gender, nil
}

func (r *genderRepository) GetGenderByName(name string) (*models.Gender, error) {
	var gender models.Gender

	db := r.db.Model(&gender)

	result := db.Debug().Where("name = ?", name).First(&gender)
	if result.Error != nil {
		return nil, result.Error
	}

	return &gender, nil
}
