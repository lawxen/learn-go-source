package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 官方文档：https://gorm.io/docs/create.html

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
	// 单表创建数据
	t, _ := time.Parse("2006-01-02", "1982-10-01")
	p1 := Program{Name: "西游记", Online: t}
	r1 := DB.Create(&p1)
	fmt.Printf("Program创建数据ID为：%v\n", p1.ID)
	fmt.Printf("返回插入记录的条数：%v\n", r1.RowsAffected)
	// 创建部门字段
	p2 := Program{Name: "红楼梦", Online: t}
	// Select是对部分字段赋值
	// 没有赋值字段如有默认值或非空类型，字段也会生成数据
	DB.Select("Name").Create(&p2)
	//DB.Omit("Online").Create(&p2)
	// 批量创建
	p3 := []Program{{Name: "三国演义"}, {Name: "水浒传"}}
	DB.Create(&p3)
	// 集合格式创建数据
	// 创建单数据
	m1 := map[string]interface{}{"Name": "封神榜"}
	DB.Model(&Program{}).Create(&m1)
	// 批量创建
	m2 := []map[string]interface{}{
		{"Name": "姜子牙"},
		{"Name": "哪吒"},
	}
	DB.Model(&Program{}).Create(&m2)
	// 创建数据关联
	u1 := User{Name: "Tom", Nationality: "China", Program: []Program{p1}}
	DB.Create(&u1)
	u2 := User{Name: "Tim", Nationality: "China", Program: []Program{{Name: "大闹天宫"}, {Name: "流浪地球"}}}
	DB.Create(&u2)
}
