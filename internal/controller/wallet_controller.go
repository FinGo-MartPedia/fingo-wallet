package controller

import (
	"net/http"

	"github.com/fingo-martpedia/fingo-wallet/constants"
	"github.com/fingo-martpedia/fingo-wallet/helpers"
	"github.com/fingo-martpedia/fingo-wallet/internal/interfaces"
	"github.com/fingo-martpedia/fingo-wallet/internal/models"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/requests"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/responses"
	"github.com/gin-gonic/gin"
)

type WalletController struct {
	WalletService interfaces.IWalletService
}

func NewWalletController(walletService interfaces.IWalletService) *WalletController {
	return &WalletController{
		WalletService: walletService,
	}
}

func (api *WalletController) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req requests.CreateWalletRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	response, err := api.WalletService.Create(c.Request.Context(), req.UserID)
	if err != nil {
		log.Error("Failed to create wallet: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, response)
}

func (api *WalletController) DebitBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req requests.DebitBalanceRequest
	)

	u, exists := c.Get("user")
	if !exists {
		log.Error("Failed to get user from context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	user, ok := u.(models.User)
	if !ok {
		log.Error("Invalid user type in context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	userID := user.ID

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	response, err := api.WalletService.DebitBalance(c.Request.Context(), int(userID), req.Amount)
	if err != nil {
		log.Error("Failed to debit balance: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, response)
}

func (api *WalletController) CreditBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req requests.DebitBalanceRequest
	)

	u, exists := c.Get("user")
	if !exists {
		log.Error("Failed to get user from context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	user, ok := u.(models.User)
	if !ok {
		log.Error("Invalid user type in context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	userID := user.ID

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	response, err := api.WalletService.CreditBalance(c.Request.Context(), int(userID), req.Amount)
	if err != nil {
		log.Error("Failed to credit balance: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, response)
}

func (api *WalletController) GetBalance(c *gin.Context) {
	var log = helpers.Logger

	u, exists := c.Get("user")
	if !exists {
		log.Error("Failed to get user from context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	user, ok := u.(models.User)
	if !ok {
		log.Error("Invalid user type in context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	userID := user.ID

	response, err := api.WalletService.GetBalance(c.Request.Context(), int(userID))
	if err != nil {
		log.Error("Failed to get balance: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, responses.BalanceResponse{Balance: response})
}

func (api *WalletController) HistoryWalletTransactions(c *gin.Context) {
	var (
		log = helpers.Logger
		req requests.WalletHistoryParam
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Error("Failed to bind request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("Failed to validate request body: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, err.Error())
		return
	}

	u, exists := c.Get("user")
	if !exists {
		log.Error("Failed to get user from context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	user, ok := u.(models.User)
	if !ok {
		log.Error("Invalid user type in context")
		helpers.SendResponse(c, http.StatusUnauthorized, constants.ErrFailedUnauthorized, nil)
		return
	}

	response, err := api.WalletService.HistoryWalletTransactions(c.Request.Context(), int(user.ID), req)
	if err != nil {
		log.Error("Failed to get wallet transactions: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrFailedServerError, err.Error())
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, response)
}
