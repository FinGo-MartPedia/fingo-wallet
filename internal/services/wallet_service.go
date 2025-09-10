package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/fingo-martpedia/fingo-wallet/internal/interfaces"
	"github.com/fingo-martpedia/fingo-wallet/internal/models"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/requests"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/responses"
	"github.com/pkg/errors"
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

func randomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(b))
}

func (s *WalletService) DebitBalance(ctx context.Context, userId int, amount float64) (responses.WalletTransactionResponse, error) {
	var response responses.WalletTransactionResponse

	wallet, err := s.WalletRepository.UpdateBalance(ctx, userId, amount)
	if err != nil {
		return response, err
	}

	reference := "WD-" + randomString(8)

	err = s.WalletRepository.CreateWalletTransaction(ctx, &models.WalletTransaction{
		WalletID:  wallet.ID,
		Amount:    amount,
		Type:      "DEBIT",
		Reference: reference,
	})
	if err != nil {
		errors.Wrap(err, "failed to create wallet transaction")
		return response, err
	}

	response.Balance = wallet.Balance
	response.Type = "DEBIT"
	response.Reference = reference

	return response, nil
}

func (s *WalletService) CreditBalance(ctx context.Context, userId int, amount float64) (responses.WalletTransactionResponse, error) {
	var response responses.WalletTransactionResponse

	existsWallet, err := s.WalletRepository.GetWalletByUserID(ctx, userId)
	if err != nil {
		return response, errors.Wrap(err, "failed to get wallet")
	}

	if existsWallet.Balance < amount {
		return response, errors.New("insufficient balance")
	}

	wallet, err := s.WalletRepository.UpdateBalance(ctx, userId, -amount)
	if err != nil {
		return response, err
	}

	reference := "WD-" + randomString(8)

	err = s.WalletRepository.CreateWalletTransaction(ctx, &models.WalletTransaction{
		WalletID:  wallet.ID,
		Amount:    amount,
		Type:      "CREDIT",
		Reference: reference,
	})
	if err != nil {
		errors.Wrap(err, "failed to create wallet transaction")
		return response, err
	}

	response.Balance = wallet.Balance
	response.Type = "CREDIT"
	response.Reference = reference

	return response, nil
}

func (s *WalletService) GetBalance(ctx context.Context, userId int) (float64, error) {
	wallet, err := s.WalletRepository.GetWalletByUserID(ctx, userId)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get wallet")
	}
	return wallet.Balance, nil
}

func (s *WalletService) HistoryWalletTransactions(ctx context.Context, userId int, param requests.WalletHistoryParam) (responses.PaginatedResult[responses.WalletTransactionResponse], error) {
	var resp responses.PaginatedResult[responses.WalletTransactionResponse]

	wallet, err := s.WalletRepository.GetWalletByUserID(ctx, userId)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get wallet")
	}

	offset := (param.Page - 1) * param.Limit
	transactions, err := s.WalletRepository.GetWalletTransactions(ctx, wallet.ID, offset, param.Limit, param.Type)
	if err != nil {
		return resp, errors.Wrap(err, "failed to get wallet transactions")
	}

	walletHistory := make([]responses.WalletTransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		walletHistory = append(walletHistory, responses.WalletTransactionResponse{
			Balance:   transaction.Amount,
			Type:      transaction.Type,
			Reference: transaction.Reference,
		})
	}

	total, err := s.WalletRepository.CountWalletTransactions(ctx, wallet.ID, param.Type)
	if err != nil {
		return resp, errors.Wrap(err, "failed to count wallet transactions")
	}

	resp = responses.PaginatedResult[responses.WalletTransactionResponse]{
		Items:      walletHistory,
		Total:      int(total),
		Page:       param.Page,
		PageSize:   param.Limit,
		TotalPages: int((total + int64(param.Limit) - 1) / int64(param.Limit)),
	}

	return resp, nil
}
