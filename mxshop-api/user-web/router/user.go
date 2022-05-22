/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 16:58:04
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 17:09:54
 */
package router

import (
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	UserRouter.GET("list", api.GetUserList)
}
