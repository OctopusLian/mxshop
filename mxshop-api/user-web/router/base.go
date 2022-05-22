/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 21:10:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:10:48
 */
package router

import (
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
