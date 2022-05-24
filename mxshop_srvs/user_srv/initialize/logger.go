/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-24 22:44:37
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-24 22:44:40
 */
package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
