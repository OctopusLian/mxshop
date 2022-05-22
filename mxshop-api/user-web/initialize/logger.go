/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 17:40:46
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 17:40:55
 */
package initialize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
