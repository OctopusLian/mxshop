package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/userop-web/api/address"
	"mxshop-api/userop-web/middlewares"
)

func InitAddressRouter(Router *gin.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)          // 轮播图列表页
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete) // 删除轮播图
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)          //新建轮播图
		AddressRouter.PUT("/:id", middlewares.JWTAuth(), address.Update)    //修改轮播图信息
	}
}
