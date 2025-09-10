package interfaces

import (
	"context"

	"github.com/fingo-martpedia/fingo-wallet/internal/models"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/requests"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/responses"
	"github.com/gin-gonic/gin"
)

type IWalletRepository interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
	UpdateBalance(ctx context.Context, userID int, amount float64) (models.Wallet, error)
	CreateWalletTransaction(ctx context.Context, walletTransaction *models.WalletTransaction) error
	GetWalletByUserID(ctx context.Context, userID int) (models.Wallet, error)
	GetWalletTransactions(ctx context.Context, walletId int, offset int, limit int, transactionType string) ([]models.WalletTransaction, error)
	CountWalletTransactions(ctx context.Context, walletId int, transactionType string) (int64, error)
}

type IWalletService interface {
	Create(ctx context.Context, userId int) (*models.Wallet, error)
	DebitBalance(ctx context.Context, userId int, amount float64) (responses.WalletTransactionResponse, error)
	CreditBalance(ctx context.Context, userId int, amount float64) (responses.WalletTransactionResponse, error)
	GetBalance(ctx context.Context, userId int) (float64, error)
	HistoryWalletTransactions(ctx context.Context, userId int, param requests.WalletHistoryParam) (responses.PaginatedResult[responses.WalletTransactionResponse], error)
}

type IWalletController interface {
	Create(c *gin.Context)
	DebitBalance(c *gin.Context)
	CreditBalance(c *gin.Context)
	GetBalance(c *gin.Context)
	HistoryWalletTransactions(c *gin.Context)
}
