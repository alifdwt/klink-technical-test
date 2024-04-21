package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Privilege string    `json:"privilege" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Member    []Member  `json:"member" gorm:"constraint:OnDelete:CASCADE"`
}
