/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-26 21:46:09
 * @LastEditors: neozhang
 * @LastEditTime: 2022-07-18 22:53:07
 */
package router

import (
	"mxshop/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha) //获取图片验证码
		BaseRouter.POST("send_sms", api.SendSms)  //短信发送
	}
}
