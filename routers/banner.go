package router

import (
	"mxshop/api"
	"mxshop/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banners").Use(middlewares.Trace())
	{
		BannerRouter.GET("", api.BannerList)                                                            // 轮播图列表页
		BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.DeleteBanner) // 删除轮播图
		BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.NewBanner)          //新建轮播图
		BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.UpdateBanner)    //修改轮播图信息
	}
}
