package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// go mod init MyGin
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁止创建数据库的外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	// 设置数据库连接池
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}
