package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://gorm.io/docs/has_many.html

// 一对多关联
type User struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(255);unique"`
	Nationality string `gorm:"type:varchar(255)"`
	Program     []Program
}

type Program struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"type:varchar(255)"`
	UserID int64
}

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 启用创建数据库的外键约束
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	// 执行数据迁移
	DB.AutoMigrate(&User{}, &Program{})
}
