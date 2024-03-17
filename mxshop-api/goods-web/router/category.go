package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/api/category"
	"mxshop-api/goods-web/middlewares"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	CategoryRouter := Router.Group("categorys").Use(middlewares.Trace())
	{
		CategoryRouter.GET("", category.List)          // 商品类别列表页
		CategoryRouter.DELETE("/:id", category.Delete) // 删除分类
		CategoryRouter.GET("/:id", category.Detail)    // 获取分类详情
		CategoryRouter.POST("", category.New)          //新建分类
		CategoryRouter.PUT("/:id", category.Update)    //修改分类信息
	}
}
