package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(192.168.0.104:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&CreditCard{})

	//在大型的系统中，我个人不建议使用外键约束，外键约束也有很大的优点： 数据的完整性
	/*
		外键约束会让给你的数据很完整，即使是业务代码有些人考虑的不严谨
		在大型的系统，高并发的系统中一般不使用外键约束，自己在业务层面保证数据的一致性
	*/
	//user := User{}
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number:    "12",
	//	UserRefer: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number:    "34",
	//	UserRefer: user.ID,
	//})
	var user User
	db.Preload("CreditCards").First(&user)
	for _, card := range user.CreditCards {
		fmt.Println(card.Number)
	}

}
