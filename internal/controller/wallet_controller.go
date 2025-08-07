package controller

import (
	"net/http"

	"github.com/fingo-martpedia/fingo-wallet/constants"
	"github.com/fingo-martpedia/fingo-wallet/helpers"
	"github.com/fingo-martpedia/fingo-wallet/internal/interfaces"
	"github.com/fingo-martpedia/fingo-wallet/internal/models/requests"
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
