package services

import (
	"context"

	"github.com/fingo-martpedia/fingo-wallet/internal/interfaces"
	"github.com/fingo-martpedia/fingo-wallet/internal/models"
)

type WalletService struct {
	WalletRepository interfaces.IWalletRepository
}

func NewWalletService(walletRepository interfaces.IWalletRepository) *WalletService {
	return &WalletService{
		WalletRepository: walletRepository,
	}
}

func (s *WalletService) Create(ctx context.Context, userId int) (*models.Wallet, error) {
	wallet := &models.Wallet{UserID: userId}
	err := s.WalletRepository.CreateWallet(ctx, wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
