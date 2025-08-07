package cmd

import (
	"log"
	"net/http"

	"github.com/fingo-martpedia/fingo-wallet/external"
	"github.com/fingo-martpedia/fingo-wallet/helpers"
	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateToken(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("authorization empty")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		log.Println("invalid authorization format")
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	accessToken := authHeader[len(bearerPrefix):]

	userData, err := external.ValidateToken(ctx.Request.Context(), accessToken)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(ctx, http.StatusUnauthorized, "unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("user", userData)

	ctx.Next()
}
