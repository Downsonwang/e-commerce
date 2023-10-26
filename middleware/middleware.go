/*
 * @Descripttion:
 * @Author:
 * @Date: 2023-07-08 13:44:33
 * @LastEditTime: 2023-08-05 14:59:12
 */
package middleware

import (
	"ecommerce-shop/tokens"
	"go/token"
	"net/http"
	"tokens"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ClientToken := ctx.Request.Header.Get("token")
		if ClientToken == "" {
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":"No authorization head"
			})
			ctx.Abort()
			return
		}
		claims,err := tokens.ValidateToken(ClientToken)
		if err != ""{
			ctx.JSON(http.StatusInternalServerError,gin.H{"error":err})
		    ctx.Abort()
			return
		}
		ctx.Set("email",claims.Email)
		ctx.Set("uid",claims.Uid)
		ctx.Next()
	}
}
