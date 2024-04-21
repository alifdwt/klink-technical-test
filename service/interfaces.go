package service

import (
	"technical-test/domain/requests/auth"
	"technical-test/domain/requests/member"
	"technical-test/domain/responses"
	"technical-test/models"
)

type AuthService interface {
	Register(input *auth.RegisterRequest) (*responses.UserResponse, error)
	Login(input *auth.LoginRequest) (*responses.Token, error)
}

type MemberService interface {
	CreateMember(userId int, genderId int, levelId int, request member.CreateMemberRequest) (*responses.MemberResponse, error)
	GetMemberAll() (*[]responses.MemberResponse, error)
	GetMemberByUserId(userId int) (*[]responses.MemberWithRelationResponse, error)
	GetMemberById(id int) (*responses.MemberWithRelationResponse, error)
}

type GenderService interface {
	CreateGender(name string) (*models.Gender, error)
	GetGenderById(id int) (*models.Gender, error)
	GetGenderByName(name string) (*models.Gender, error)
}

type LevelService interface {
	CreateLevel(name string) (*models.Level, error)
	GetLevelById(id int) (*models.Level, error)
	GetLevelByName(name string) (*models.Level, error)
}
