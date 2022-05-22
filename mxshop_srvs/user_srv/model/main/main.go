/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-21 22:35:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 16:20:18
 */
package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mxshop_srvs/user_srv/model"
	"os"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //创建的表末尾不加s
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodePwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512%s$%s", salt, encodePwd)
	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("bobby%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}
}
