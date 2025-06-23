package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://gorm.io/docs/has_one.html
// https://gorm.io/docs/belongs_to.html

// 一对一关联实现方式一
type User struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(255);unique"`
	Nationality string `gorm:"type:varchar(255)"`
	Masterpiece string `gorm:"type:varchar(255)"`
}

type Info struct {
	ID     uint   `gorm:"primarykey"`
	Birth  string `gorm:"type:varchar(255)"`
	Elapse string `gorm:"type:varchar(255)"`
	UserID int64
	User   User
}

// 一对一关联实现方式二
type User1 struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(255);unique"`
	Nationality string `gorm:"type:varchar(255)"`
	Masterpiece string `gorm:"type:varchar(255)"`
	Info1       Info1
}

type Info1 struct {
	ID      uint   `gorm:"primarykey"`
	Birth   string `gorm:"type:varchar(255)"`
	Elapse  string `gorm:"type:varchar(255)"`
	User1ID int64
}

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 启用创建数据库的外键约束
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	// 执行数据迁移
	DB.AutoMigrate(&User{}, &Info{}, &User1{}, &Info1{})
}
