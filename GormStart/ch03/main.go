package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	a := []int{1, 2, 3}
	b := a[:]
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
	_ = db.AutoMigrate(&User{}) //此处应该有sql语句
	user := User{
		Name: "bobby2",
	}
	fmt.Println(user.ID)
	result := db.Create(&user)
	fmt.Println(user.ID)             // 返回插入数据的主键
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	//db.Model(&User{ID:1}).Update("Name", "")
	//updates语句不会更新零值，但是update语句会更新
	//empty := ""
	//db.Model(&User{ID:1}).Updates(User{Email: &empty})
	//解决仅更新非零值字段的方法有两种
	/*
		1. 将string 设置为 *string
		2. 使用sql的NULLxxx来解决
	*/

}
