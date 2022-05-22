/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 21:56:10
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:56:20
 */
package global

import (
	"mxshop_srvs/goods_srv/config"

	"gorm.io/gorm"

	"github.com/olivere/elastic/v7"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig

	EsClient *elastic.Client
)
