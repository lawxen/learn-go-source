package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// https://gorm.io/docs/models.html
// https://gorm.io/docs/data_types.html
// https://gorm.io/docs/indexes.html
// https://gorm.io/docs/migration.html

type Base struct {
	gorm.Model
	// 创建唯一索引
	Username string `gorm:"type:varchar(255);unique"`
	Password string `gorm:"type:varchar(255)"`
}

type User struct {
	Base Base `gorm:"embedded"`
	// 登录时间，autoUpdateTime是自动更新
	LoginTime time.Time `gorm:"autoUpdateTime:milli"`
}

type Info struct {
	gorm.Model
	UserId int64 `json:"userId"`
	// 创建外键关联
	User User `gorm:"foreignkey:UserId"`
	// 创建普通索引
	Age int64 `gorm:"index:,class:,type:,comment:,sort:"`
}

// 结构体User默认的数据表名为Users
// 如果自定义数据表名，可自定义TableName方法
func (User) TableName() string {
	return "myUser"
}

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁止创建数据库的外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	// 执行数据迁移
	DB.AutoMigrate(&User{}, &Info{})
}
