package middlewares

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/userop-web/models"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {
	//将一些共用的代码抽出来然后共用 - 版本管理
	//如果不抽出来
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}

}
