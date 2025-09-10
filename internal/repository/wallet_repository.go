package repository

import (
	"context"
	"fmt"

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
	return r.DB.Create(wallet).Error
}

func (r *WalletRepository) UpdateBalance(ctx context.Context, userID int, amount float64) (models.Wallet, error) {
	var wallet models.Wallet
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Raw("SELECT id, user_id, balance FROM wallets WHERE user_id = ? FOR UPDATE", userID).Scan(&wallet).Error
		if err != nil {
			return err
		}

		if wallet.ID == 0 {
			return gorm.ErrRecordNotFound
		}

		if (wallet.Balance + amount) < 0 {
			return fmt.Errorf("balance is negative")
		}

		err = tx.Exec("UPDATE wallets SET balance = balance + ? WHERE user_id = ?", amount, userID).Error
		if err != nil {
			return err
		}

		return tx.Raw("SELECT id, user_id, balance FROM wallets WHERE user_id = ?", userID).Scan(&wallet).Error
	})
	return wallet, err
}

func (r *WalletRepository) CreateWalletTransaction(ctx context.Context, walletTransaction *models.WalletTransaction) error {
	return r.DB.Create(walletTransaction).Error
}

func (r *WalletRepository) GetWalletByUserID(ctx context.Context, userID int) (models.Wallet, error) {
	var wallet models.Wallet
	err := r.DB.Where("user_id = ?", userID).First(&wallet).Error
	return wallet, err
}

func (r *WalletRepository) GetWalletTransactions(ctx context.Context, walletId int, offset int, limit int, transactionType string) ([]models.WalletTransaction, error) {
	var walletTransactions []models.WalletTransaction
	sql := r.DB
	if transactionType != "" {
		sql = sql.Where("Type = ?", transactionType)
	}
	err := sql.Limit(limit).Offset(offset).Order("id desc").Find(&walletTransactions, "wallet_id = ?", walletId).Error
	return walletTransactions, err
}

func (r *WalletRepository) CountWalletTransactions(ctx context.Context, walletId int, transactionType string) (int64, error) {
	var total int64
	sql := r.DB.Model(&models.WalletTransaction{}).Where("wallet_id = ?", walletId)
	if transactionType != "" {
		sql = sql.Where("Type = ?", transactionType)
	}
	err := sql.Count(&total).Error
	return total, err
}
