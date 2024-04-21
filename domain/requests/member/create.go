package member

import (
	"github.com/go-playground/validator/v10"
)

type CreateMemberRequest struct {
	NamaDepan    string `json:"nama_depan" validate:"required"`
	NamaBelakang string `json:"nama_belakang" validate:"required"`
	Birthdate    string `json:"birthdate" validate:"required"`
	Gender       string `json:"gender" validate:"required,oneof=Male Female"`
	Level        string `json:"level" validate:"required"`
}

func (c *CreateMemberRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
