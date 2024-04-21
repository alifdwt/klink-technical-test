package repository

import (
	"technical-test/domain/requests/member"
	"technical-test/models"
	"time"

	"gorm.io/gorm"
)

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{db}
}

func (r *memberRepository) CreateMember(userId int, genderId int, levelId int, request member.CreateMemberRequest) (*models.Member, error) {
	var member models.Member

	db := r.db.Model(&member)

	member.UserID = userId
	member.GenderID = genderId
	member.LevelID = levelId

	member.NamaDepan = request.NamaDepan
	member.NamaBelakang = request.NamaBelakang

	birthdate, err := time.Parse("2006-01-02", request.Birthdate)
	if err != nil {
		return nil, err
	}
	member.Birthdate = birthdate

	member.JoinDate = time.Now()

	result := db.Debug().Create(&member)
	if result.Error != nil {
		return nil, result.Error
	}

	return &member, nil
}

func (r *memberRepository) GetMemberAll() (*[]models.Member, error) {
	var member []models.Member

	db := r.db.Model(&member)

	result := db.Debug().Preload("User").Preload("Gender").Preload("Level").Find(&member)
	if result.Error != nil {
		return nil, result.Error
	}

	return &member, nil
}

func (r *memberRepository) GetMemberByUserId(userId int) (*[]models.Member, error) {
	var member []models.Member

	db := r.db.Model(&member)

	result := db.Debug().Preload("User").Preload("Gender").Preload("Level").Where("user_id = ?", userId).Find(&member)
	if result.Error != nil {
		return nil, result.Error
	}

	return &member, nil
}

func (r *memberRepository) GetMemberById(id int) (*models.Member, error) {
	var member models.Member

	db := r.db.Model(&member)

	result := db.Debug().Preload("User").Preload("Gender").Preload("Level").Where("id = ?", id).First(&member)
	if result.Error != nil {
		return nil, result.Error
	}

	return &member, nil
}
