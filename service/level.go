package service

import (
	"technical-test/models"
	"technical-test/pkg/logger"
	"technical-test/repository"
)

type levelService struct {
	repository repository.LevelRepository
	log        logger.Logger
}

func NewLevelService(level repository.LevelRepository, log logger.Logger) *levelService {
	return &levelService{
		repository: level,
		log:        log,
	}
}

func (s *levelService) CreateLevel(name string) (*models.Level, error) {
	res, err := s.repository.CreateLevel(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *levelService) GetLevelById(id int) (*models.Level, error) {
	res, err := s.repository.GetLevelById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *levelService) GetLevelByName(name string) (*models.Level, error) {
	res, err := s.repository.GetLevelByName(name)
	if err != nil {
		return nil, err
	}

	return res, nil
}
