package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// 官方文档：https://gorm.io/docs/query.html

// 多对多关联
type User struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"type:varchar(255);unique"`
	Nationality string    `gorm:"type:varchar(255)"`
	Program     []Program `gorm:"many2many:user_program"`
	Deleted     gorm.DeletedAt
}

type Program struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(255)"`
	//Online time.Time `gorm:"null"`
	Online  time.Time `gorm:"autoCreateTime"`
	Deleted gorm.DeletedAt
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

	p1 := Program{}
	// 查询数据表第一条数据，默认以主键升序排序
	// SQL语句：SELECT * FROM programs ORDER BY id LIMIT 1;
	DB.First(&p1)
	// 获取最后一条记录
	// SQL语句：SELECT * FROM programs ORDER BY id DESC LIMIT 1;
	res := DB.Last(&p1)
	// 返回找到的记录数
	fmt.Printf("返回找到的记录数：%v\n", res.RowsAffected)
	// 返回error或者nil
	fmt.Printf("返回异常信息：%v\n", res.Error)

	// 查询所有对象
	u2 := []User{}
	// 如果存在外键属性
	// 调用Preload(clause.Associations)获取外键数据
	// 参考官方文档：https://gorm.io/docs/preload.html
	DB.Preload(clause.Associations).Find(&u2)
	// 设置查询条件，默认以主键筛选
	DB.Preload(clause.Associations).Find(&u2, []int{1})
	// 查询非主键字段
	DB.Preload(clause.Associations).Find(&u2, "name IN ?", []string{"Tim"})
	// 等同于
	DB.Preload(clause.Associations).Where("name IN ?", []string{"Tim"}).Find(&u2)
}
