package repository

import (
	"technical-test/domain/requests/auth"
	"technical-test/models"
	"technical-test/pkg/dotenv"

	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
	config dotenv.Config
}

func NewUserRepository(db *gorm.DB, config dotenv.Config) *userRepository {
	return &userRepository{db, config}
}

func (r *userRepository) CreateUser(registerReq *auth.RegisterRequest) (*models.User, error) {
	var user models.User

	user.Username = registerReq.Username
	user.Password = registerReq.Password

	if registerReq.AdminCode == r.config.JWT_SECRET {
		user.Privilege = "admin"
	} else {
		user.Privilege = "member"
	}

	db := r.db.Model(&user)

	result := db.Debug().Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) GetUserById(id int) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUserById := db.Debug().Preload("Member").Where("id = ?", id).First(&user)
	if checkUserById.Error != nil {
		return nil, checkUserById.Error
	}

	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	db := r.db.Model(user)

	checkUserByUsername := db.Debug().Preload("Member").Where("username = ?", username).First(&user)
	if checkUserByUsername.Error != nil {
		return nil, checkUserByUsername.Error
	}

	return &user, nil
}
