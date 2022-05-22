/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 16:43:24
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:44:00
 */
package main

import (
	"fmt"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	myvalidator "mxshop-api/user-web/validator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	//初始化logger
	initialize.InitLogger()
	//初始化config
	initialize.InitConfig()
	//初始化routers
	Router := initialize.Routers()
	//初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", myvalidator.ValidateMobile)
	}
	/*
		1. S()可以获取一个全局的sugar，可以让我们自己设置一个全局的logger
		2. 日志是分级别的，debug， info ， warn， error， fetal
		3. S函数和L函数很有用， 提供了一个全局的安全访问logger的途径
	*/
	zap.S().Debugf("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败:", err.Error())
	}
}
