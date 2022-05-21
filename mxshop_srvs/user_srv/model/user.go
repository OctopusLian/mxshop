/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-21 22:23:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-21 23:42:19
 */
package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int32     `gorm:"primarykey"`
	CreateAt  time.Time `gorm:"column:add_time"`
	UpdateAt  time.Time `gorm:"column:update_time"`
	DeleteAt  gorm.DeletedAt
	IsDeleted bool `gorm:"column:is_deleted"`
}

/*
1. 密文 2. 密文不可反解
	1. 对称加密，不满足 密文不可反解的条件，不考虑
	2. 非对称加密
	3. md5 信息摘要算法
	密码如果不可以反解，用户找回密码
*/
type User struct {
	BaseModel
	//要加索引，不然后期数据量大，查询会很慢
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	NickName string `gorm:"type:varchar(20)"`
	//加指针，是为了防止Gorm保存的时候出错
	Birthday *time.Time `gorm:"type:datetime"`
	Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'femail表示女，male表示男'"`
	Role     int        `gorm:"column:role;default:1;type:int comment '1表示普通用户，2表示管理员'"`
}
