package repository

import (
	"context"

	"github.com/fingo-martpedia/fingo-wallet/internal/models"
	"gorm.io/gorm"
)

type WalletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		DB: db,
	}
}

func (r *WalletRepository) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	return r.DB.Create(&wallet).Error
}
