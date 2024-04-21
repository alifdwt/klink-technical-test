package responses

import "time"

type MemberResponse struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	GenderID     int       `json:"gender_id"`
	LevelID      int       `json:"level_id"`
	NamaDepan    string    `json:"nama_depan"`
	NamaBelakang string    `json:"nama_belakang"`
	Birthdate    time.Time `json:"birthdate"`
	JoinDate     time.Time `json:"join_date"`
	UpdateAt     time.Time `json:"update_at"`
}

type MemberWithRelationResponse struct {
	ID           int            `json:"id"`
	UserID       int            `json:"user_id"`
	GenderID     int            `json:"gender_id"`
	LevelID      int            `json:"level_id"`
	NamaDepan    string         `json:"nama_depan"`
	NamaBelakang string         `json:"nama_belakang"`
	Birthdate    time.Time      `json:"birthdate"`
	JoinDate     time.Time      `json:"join_date"`
	UpdateAt     time.Time      `json:"update_at"`
	User         UserResponse   `json:"user"`
	Gender       GenderResponse `json:"gender"`
	Level        LevelResponse  `json:"level"`
}
