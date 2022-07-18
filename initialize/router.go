/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-07-18 22:50:23
 * @LastEditors: neozhang
 * @LastEditTime: 2022-07-18 22:51:51
 */
package initialize

import (
	"mxshop/middlewares"
	"mxshop/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouters() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}
