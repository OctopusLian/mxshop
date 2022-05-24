/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-21 23:21:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-24 22:44:14
 */
package global

import (
	"mxshop_srvs/goods_srv/config"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	//NacosConfig  config.NacosConfig
)
