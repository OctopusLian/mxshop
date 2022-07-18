/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-26 23:00:30
 * @LastEditors: neozhang
 * @LastEditTime: 2022-07-18 22:53:12
 */
package router

import (
	"mxshop/api"
	"mxshop/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)

		UserRouter.GET("detail", middlewares.JWTAuth(), api.GetUserDetail)
		UserRouter.PATCH("update", middlewares.JWTAuth(), api.UpdateUser)
	}
}
