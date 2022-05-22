/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 17:12:38
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 17:44:39
 */
package initialize

import (
	router "mxshop-api/user-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)

	return Router
}
