package service

import (
	"errors"
	"technical-test/domain/requests/auth"
	"technical-test/domain/responses"
	"technical-test/mapper"
	"technical-test/pkg/dotenv"
	"technical-test/pkg/hashing"
	"technical-test/pkg/logger"
	"technical-test/pkg/token"
	"technical-test/repository"
)

type authService struct {
	config dotenv.Config
	user   repository.UserRepository
	hash   hashing.Hashing
	log    logger.Logger
	token  token.Maker
	mapper mapper.UserMapping
}

func NewAuthService(config dotenv.Config, user repository.UserRepository, hash hashing.Hashing, log logger.Logger, token token.Maker, mapper mapper.UserMapping) *authService {
	return &authService{
		config: config,
		user:   user,
		hash:   hash,
		log:    log,
		token:  token,
		mapper: mapper,
	}
}

func (s *authService) Register(input *auth.RegisterRequest) (*responses.UserResponse, error) {
	hashing, err := s.hash.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	input.Password = hashing

	user, err := s.user.CreateUser(input)
	if err != nil {
		return nil, err
	}

	mapper := s.mapper.ToUserResponse(user)

	return mapper, nil
}

func (s *authService) Login(input *auth.LoginRequest) (*responses.Token, error) {
	res, err := s.user.GetUserByUsername(input.Username)
	if err != nil {
		return nil, errors.New("error while get user by username: " + err.Error())
	}

	err = s.hash.ComparePassword(res.Password, input.Password)
	if err != nil {
		return nil, errors.New("error while compare password: " + err.Error())
	}

	accessToken, err := s.createAccessToken(res.ID, res.Username, res.Privilege)
	if err != nil {
		return nil, err
	}

	return &responses.Token{
		Token: accessToken,
	}, nil

}

func (s *authService) createAccessToken(id int, username string, privilege string) (string, error) {
	res, err := s.token.CreateToken(id, username, privilege, "access", s.config.ACCESS_TOKEN_DURATION)
	if err != nil {
		return "", err
	}

	return res, nil
}
