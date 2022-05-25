package global

import (
	"mxshop/config"

	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	Trans        ut.Translator
)
