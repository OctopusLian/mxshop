/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-21 23:21:55
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-21 23:24:25
 */
package global

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
	// ServerConfig config.ServerConfig
	// NacosConfig  config.NacosConfig
)

func init() {
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io writer
		logger.Config{
			SlowThreshold: time.Second, //慢 SQL 阀值
			LogLevel:      logger.Info,
			Colorful:      true, //禁用彩色打印
		},
	)
	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //创建的表末尾不加s
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// _ = db.AutoMigrate(&model.User{})
}
