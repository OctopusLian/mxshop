package routers

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
