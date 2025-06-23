package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 官方文档：https://gorm.io/docs/transactions.html

type User struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(255);unique"`
	Nationality string `gorm:"type:varchar(255)"`
}

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 全局关闭事务
		SkipDefaultTransaction: false,
	})
	// 执行数据迁移
	DB.AutoMigrate(&User{})
	// 创建新会话关闭事务
	//tx := DB.Session(&gorm.Session{SkipDefaultTransaction: true})
	//u := User{}
	//tx.First(&u)

	// 自动事务
	DB.Transaction(func(tx *gorm.DB) error {
		// 执行数据操作
		if err := tx.Create(&User{Name: "Tom"}).Error; err != nil {
			fmt.Printf("事务异常1，回滚了\n")
			return err
		}
		fmt.Printf("事务正常1，提交了\n")
		return nil
	})

	// 手动事务
	// 创建事务对象tx
	tx := DB.Begin()
	// 执行数据操作
	err := tx.Create(&User{Name: "Tom"}).Error
	if err != nil {
		// 事务回滚
		fmt.Printf("事务异常2，回滚了\n")
		// 回滚整个事务
		tx.Rollback()
	} else {
		// 事务提交
		fmt.Printf("事务正常2，提交了\n")
		tx.Commit()
	}

	// 事务的保存点和回滚至保存点
	tx1 := DB.Begin()
	tx1.Create(&User{Name: "Lucy"})
	// 设置保存点
	tx1.SavePoint("sp1")
	tx1.Create(&User{Name: "Betty"})
	// 回滚到指定保存点
	tx1.RollbackTo("sp1")
	tx1.Commit()
}
