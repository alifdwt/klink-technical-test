package repository

import (
	"technical-test/pkg/dotenv"

	"gorm.io/gorm"
)

type Repositories struct {
	User   UserRepository
	Member MemberRepository
	Gender GenderRepository
	Level  LevelRepository
	Config dotenv.Config
}

func NewRepository(db *gorm.DB, config dotenv.Config) *Repositories {
	return &Repositories{
		User:   NewUserRepository(db, config),
		Member: NewMemberRepository(db),
		Gender: NewGenderRepository(db),
		Level:  NewLevelRepository(db),
	}
}
