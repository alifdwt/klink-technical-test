package repository

import (
	"technical-test/domain/requests/auth"
	"technical-test/domain/requests/member"
	"technical-test/models"
)

type UserRepository interface {
	CreateUser(registerReq *auth.RegisterRequest) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type MemberRepository interface {
	CreateMember(userId int, genderId int, levelId int, request member.CreateMemberRequest) (*models.Member, error)
	GetMemberAll() (*[]models.Member, error)
	GetMemberByUserId(userId int) (*[]models.Member, error)
	GetMemberById(id int) (*models.Member, error)
}

type GenderRepository interface {
	CreateGender(name string) (*models.Gender, error)
	GetGenderById(id int) (*models.Gender, error)
	GetGenderByName(name string) (*models.Gender, error)
}

type LevelRepository interface {
	CreateLevel(name string) (*models.Level, error)
	GetLevelById(id int) (*models.Level, error)
	GetLevelByName(name string) (*models.Level, error)
}
