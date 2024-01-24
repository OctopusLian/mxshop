package main

import (
	"database/sql"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Language struct {
	gorm.Model
	Name    string
	AddTime sql.NullTime //每个记录创建的时候自动加上当前时间加入到AddTime中
}

//func (l *Language) BeforeCreate(tx *gorm.DB) (err error){
//	l.AddTime = time.Now()
//	return
//}

//在gorm中可以通过给某一个struct添加TableName方法来自定义表名
//func (Language) TableName() string{
//	return "my_language"
//}

/*
1. 我们自己定义表名是什么
2. 统一的给所有的表名加上一个前缀
*/
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
	//NamingStrategy和Tablename不能同时配置，
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshop_",
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "python",
	})
}
