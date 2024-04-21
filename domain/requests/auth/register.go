package auth

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	AdminCode string `json:"admin_code"`
}

func (l *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
