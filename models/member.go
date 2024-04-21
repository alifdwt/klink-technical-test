package models

import "time"

type Member struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	UserID       int       `json:"user_id" gorm:"not null"`
	GenderID     int       `json:"gender_id" gorm:"not null"`
	LevelID      int       `json:"level_id" gorm:"not null"`
	NamaDepan    string    `json:"nama_depan" gorm:"not null"`
	NamaBelakang string    `json:"nama_belakang" gorm:"not null"`
	Birthdate    time.Time `json:"birthdate" gorm:"not null"`
	JoinDate     time.Time `json:"join_date" gorm:"not null"`
	UpdatedAt    time.Time `json:"updated_at"`
	User         User      `json:"user"`
	Gender       Gender    `json:"gender"`
	Level        Level     `json:"level"`
}
