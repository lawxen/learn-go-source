package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 官方文档：https://gorm.io/docs/hooks.html

type User struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(255);unique"`
	Nationality string `gorm:"type:varchar(255)"`
}

// 自定义创建数据前的钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("创建数据之前\n")
	return
}

// 自定义创建数据后的钩子函数
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Printf("创建数据之后\n")
	return
}

// 自定义更新数据前的钩子函数
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Printf("更新数据之前\n")
	return
}

// 自定义更新数据后的钩子函数
func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Printf("更新数据之后\n")
	return
}

// 自定义删除数据前的钩子函数
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("删除数据之前\n")
	return
}

// 自定义删除数据后的钩子函数
func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Printf("删除数据之后\n")
	return
}

// 自定义查询数据后的钩子函数
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	fmt.Printf("查询数据之后\n")
	return
}

func main() {
	// 定义数据库连接对象
	var dsn = fmt.Sprintf("root:1234@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local")
	var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 执行数据迁移
	DB.AutoMigrate(&User{})
	// 创建数据
	DB.Create(&User{Name: "Mary", Nationality: "USA"})
	// 更新数据
	DB.Model(&User{}).Where("name = ?", "Mary").Update("name", "Lucy")
	// 删除数据
	DB.Delete(&User{}, "name = ?", "Lucy")
	// 查询数据
	u := User{}
	DB.First(&u)
}
