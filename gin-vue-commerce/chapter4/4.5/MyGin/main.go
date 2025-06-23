package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 官方文档：https://gorm.io/docs/update.html

// 多对多关联
type User struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);unique"`
	Nationality string    `gorm:"type:varchar(255)"`
	Program     []Program `gorm:"many2many:user_program"`
}

type Program struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(255)"`
	//Online time.Time `gorm:"null"`
	Online time.Time `gorm:"autoCreateTime"`
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
	// 数据根据主键信息执行创建或更新
	// p1的属性ID存在数据库，执行数据更新
	p1 := Program{ID: 1, Name: "西游记", Online: time.Now()}
	DB.Save(&p1)
	// p2的属性ID不存在数据库或没有没有赋值，执行数据创建
	p2 := Program{Name: "西游记", Online: time.Now()}
	//p2 := Program{ID: 10000, Name: "西游记", Online: time.Now()}
	DB.Save(&p2)
	// 更新单个字段
	// 查找数据表某行数据，再调用Update()更新某行数据
	// 如不设置查询条件，则视为全表更新
	DB.Model(&Program{}).Where("id = ?", 2).Update("online", time.Now())
	// 更新多个字段
	// 查找数据表某行数据，再调用Updates()更新某行数据
	// 如不设置查询条件，则视为全表更新
	// 使用结构体对象更新
	DB.Model(&Program{}).Where("id = ?", 3).Updates(Program{Name: "三国A"})
	// 使用集合更新
	DB.Model(&Program{}).Where("id = ?", 4).Updates(map[string]interface{}{"name": "水浒传A", "online": time.Now()})
	// 更新单个字段，不执行钩子函数
	DB.Model(&Program{}).Where("id = ?", 5).UpdateColumn("online", time.Now())
	// 更新多个字段，不执行钩子函数
	DB.Model(&Program{}).Where("id = ?", 6).UpdateColumns(Program{Name: "姜子牙A", Online: time.Now()})
}
