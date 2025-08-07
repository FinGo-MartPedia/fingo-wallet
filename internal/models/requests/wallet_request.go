package requests

import "github.com/go-playground/validator/v10"

type CreateWalletRequest struct {
	UserID int `json:"user_id" validate:"required"`
}

func (l CreateWalletRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
