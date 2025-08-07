package interfaces

import (
	"context"

	"github.com/fingo-martpedia/fingo-wallet/internal/models"
	"github.com/gin-gonic/gin"
)

type IWalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletService interface {
	Create(ctx context.Context, userId int) (*models.Wallet, error)
}

type IWalletController interface {
	Create(c *gin.Context)
}
