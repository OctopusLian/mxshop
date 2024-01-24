package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint
	MyName       string `gorm:"column:name"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type BaseModel struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"` //为什么使用int32， bigint
	CreatedAt time.Time      `gorm:"column:add_time" json:"-"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	IsDeleted bool           `json:"-"`
}

type OrderInfo struct {
	BaseModel

	User    int32  `gorm:"type:int;index"`
	OrderSn string `gorm:"type:varchar(30);index"` //订单号，我们平台自己生成的订单号
	PayType string `gorm:"type:varchar(20) comment 'alipay(支付宝)， wechat(微信)'"`

	//status大家可以考虑使用iota来做
	Status     string `gorm:"type:varchar(20)  comment 'PAYING(待支付), TRADE_SUCCESS(成功)， TRADE_CLOSED(超时关闭), WAIT_BUYER_PAY(交易创建), TRADE_FINISHED(交易结束)'"`
	TradeNo    string `gorm:"type:varchar(100) comment '交易号'"` //交易号就是支付宝的订单号 查账
	OrderMount float32
	PayTime    *time.Time `gorm:"type:datetime"`

	Address      string `gorm:"type:varchar(100)"`
	SignerName   string `gorm:"type:varchar(20)"`
	SingerMobile string `gorm:"type:varchar(11)"`
	Post         string `gorm:"type:varchar(20)"`
}

type OrderGoods struct {
	BaseModel

	Order int32 `gorm:"type:int;index"`
	Goods int32 `gorm:"type:int;index"`

	//把商品的信息保存下来了 ， 字段冗余， 高并发系统中我们一般都不会遵循三范式  做镜像 记录
	GoodsName  string `gorm:"type:varchar(100);index"`
	GoodsImage string `gorm:"type:varchar(200)"`
	GoodsPrice float32
	Nums       int32 `gorm:"type:int"`
}

func GenerateOrderSn(userId int32) string {
	//订单号的生成规则
	/*
		年月日时分秒+用户id+2位随机数
	*/
	now := time.Now()
	rand.Seed(time.Now().UnixNano())
	orderSn := fmt.Sprintf("%d%d%d%d%d%d%d%d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Nanosecond(),
		userId, rand.Intn(90)+10,
	)
	return orderSn
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
	//通过where查询
	//var user User
	//var users []User
	//db.Where("name = ?", "bobby").First(&user)
	//db.Where(&User{MyName:"bobby"}).First(&user)
	//db.Where(&User{MyName:"bobby1", Age: 0}).Find(&users)
	//db.Where(map[string]interface{}{"name": "bobby", "age":0}).Find(&users)
	//for _, user := range users{
	//	fmt.Println(user.ID)
	//}
	_ = db.AutoMigrate(OrderInfo{}, OrderGoods{})

	tx := db.Begin()
	order := OrderInfo{
		OrderSn:      GenerateOrderSn(1),
		OrderMount:   44.2,
		Address:      "北京市",
		SignerName:   "bobby",
		SingerMobile: "18787878787",
		Post:         "请尽快发货",
	}
	if result := tx.Save(&order); result.RowsAffected == 0 {
		tx.Rollback()
		fmt.Println("创建订单失败")
	}

	var orderGoods []*OrderGoods

	orderGoods = append(orderGoods, &OrderGoods{
		Order:      order.ID,
		Goods:      421,
		GoodsName:  "烟台红富士苹果12个 净重2.6kg以上 单果190-240g 新生鲜水果",
		GoodsImage: "https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/df392d01993cdab9de740fe17798bda1",
		GoodsPrice: 44.9,
		Nums:       1,
	})

	for _, orderGood := range orderGoods {
		orderGood.Order = order.ID
	}

	//批量插入orderGoods
	if result := tx.CreateInBatches(orderGoods, 100); result.RowsAffected == 0 {
		tx.Rollback()
		fmt.Println("创建订单失败")
	}
	tx.Commit()

	//查询方式条件有三种 1. string 2. struct 3. map

}
