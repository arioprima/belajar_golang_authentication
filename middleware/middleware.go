package middleware

import (
	"github.com/arioprima/belajar_golang_authentication/config"
	"github.com/arioprima/belajar_golang_authentication/repository"
	"github.com/arioprima/belajar_golang_authentication/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Deserialization(repository repository.CustomersRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.GetHeader("Authorization") // Use GetHeader instead of Request.Header.Get
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 || fields[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		token = fields[1]

		config, err := config.LoadConfig()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		userID := sub.(string) // Assuming sub is a string UUID
		result, err := repository.FindById(ctx, nil, userID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		ctx.Set("customer", result.Username) // Use "customer" instead of "customers"
		ctx.Next()
	}
}
