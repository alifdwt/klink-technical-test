package models

type Level struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}
