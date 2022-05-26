/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-25 23:00:19
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-26 22:52:58
 */
package global

import (
	"mxshop/config"

	ut "github.com/go-playground/universal-translator"
	"github.com/olivere/elastic"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	Trans        ut.Translator
	EsClient     *elastic.Client
)
