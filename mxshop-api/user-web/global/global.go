/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 18:14:16
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:18:32
 */
package global

import (
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans         ut.Translator
	UserSrvClient proto.UserClient
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	NacosConfig   *config.NacosConfig  = &config.NacosConfig{}
)
