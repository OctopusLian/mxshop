package global

import (
	ut "github.com/go-playground/universal-translator"
	"mxshop-api/oss-web/config"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
