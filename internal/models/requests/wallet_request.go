package requests

import "github.com/go-playground/validator/v10"

type CreateWalletRequest struct {
	UserID int `json:"user_id" validate:"required"`
}

func (l CreateWalletRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type DebitBalanceRequest struct {
	Amount float64 `json:"amount" validate:"required,gte=0"`
}

func (l DebitBalanceRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type WalletHistoryParam struct {
	Page  int    `form:"page" validate:"required,gte=1"`
	Limit int    `form:"limit" validate:"required,gte=1"`
	Type  string `form:"type"`
}

func (l WalletHistoryParam) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
