package service

import (
	"technical-test/mapper"
	"technical-test/pkg/dotenv"
	"technical-test/pkg/hashing"
	"technical-test/pkg/logger"
	"technical-test/pkg/token"
	"technical-test/repository"
)

type Service struct {
	Auth   AuthService
	Member MemberService
	Gender GenderService
	Level  LevelService
}

type Deps struct {
	Config     dotenv.Config
	Repository *repository.Repositories
	Hashing    hashing.Hashing
	TokenMaker token.Maker
	Logger     logger.Logger
	Mapper     mapper.Mapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth:   NewAuthService(deps.Config, deps.Repository.User, deps.Hashing, deps.Logger, deps.TokenMaker, deps.Mapper.UserMapper),
		Member: NewMemberService(deps.Repository.Member, deps.Logger, deps.Mapper.MemberMapper),
		Gender: NewGenderService(deps.Repository.Gender, deps.Logger),
		Level:  NewLevelService(deps.Repository.Level, deps.Logger),
	}
}
