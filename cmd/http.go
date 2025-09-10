package cmd

import (
	"log"

	"github.com/fingo-martpedia/fingo-wallet/helpers"
	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependencies := InitDependency()

	r := gin.Default()

	apiV1 := r.Group("/api/v1/wallet")
	apiV1.POST("/", dependencies.WalletController.Create)
	apiV1.POST("/debit", dependencies.MiddlewareValidateToken, dependencies.WalletController.DebitBalance)
	apiV1.POST("/credit", dependencies.MiddlewareValidateToken, dependencies.WalletController.CreditBalance)
	apiV1.GET("/balance", dependencies.MiddlewareValidateToken, dependencies.WalletController.GetBalance)
	apiV1.GET("/history", dependencies.MiddlewareValidateToken, dependencies.WalletController.HistoryWalletTransactions)

	err := r.Run(":" + helpers.GetEnv("PORT", "8081"))
	if err != nil {
		log.Fatal(err)
	}
}
