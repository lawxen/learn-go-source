package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 官方文档：https://gorm.io/docs/sql_builder.html

// 多对多关联
type User struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);unique"`
	Nationality string    `gorm:"type:varchar(255)"`
	Program     []Program `gorm:"many2many:user_program"`
}

type Program struct {
	ID     uint      `gorm:"primarykey"`
	Name   string    `gorm:"type:varchar(255)"`
	Online time.Time `gorm:"autoCreateTime"`
}

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 启用创建数据库的外键约束
		DisableForeignKeyConstraintWhenMigrating: false,
		// 是否允许全表删除
		AllowGlobalUpdate: false,
	})
	// 执行数据迁移
	DB.AutoMigrate(&User{}, &Program{})
	// 使用Raw()+Scan()执行原生SQL
	u1 := User{}
	DB.Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&u1)
	fmt.Printf("数据表users的id=1的数据：%v\n", u1)
	// Exex()执行原生SQL，但无法返回执行结果
	DB.Exec("UPDATE programs SET `name` = ? WHERE id IN ?", "大话西游", []int{10})
}
