package service

import (
	"technical-test/models"
	"technical-test/pkg/logger"
	"technical-test/repository"
)

type genderService struct {
	repository repository.GenderRepository
	log        logger.Logger
}

func NewGenderService(gender repository.GenderRepository, log logger.Logger) *genderService {
	return &genderService{
		repository: gender,
		log:        log,
	}
}

func (s *genderService) CreateGender(name string) (*models.Gender, error) {
	res, err := s.repository.CreateGender(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *genderService) GetGenderById(id int) (*models.Gender, error) {
	res, err := s.repository.GetGenderById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *genderService) GetGenderByName(name string) (*models.Gender, error) {
	res, err := s.repository.GetGenderByName(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}
