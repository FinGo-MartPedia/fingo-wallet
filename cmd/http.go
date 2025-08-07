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

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
