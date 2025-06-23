package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// 官方文档：https://gorm.io/docs/delete.html

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

	// 软删除不会删除表数据，在数据中记录删除时间
	// 必须数据表定义了gorm.DeletedAt类型字段，用于记录删除时间
	// 删除全表数据以及外键关联
	// 调用Select(clause.Associations)同时删除关联
	// 如果在数据表Program删除数据，但与User存在关联
	// 则无法删除成功，因为结构体Program没有外键属性
	u := []User{}
	DB.Find(&u).Select(clause.Associations).Delete(&u)
	// 删除符合条件的所有数据
	// 默认按主键删除
	DB.Delete(&Program{ID: 1})
	// 如果参数conds是整型或字符串，代表删除一行
	// 如果参数conds是切片，代表批量删除
	//DB.Delete(&Program{}, []int{9, 10})
	// 等价于
	DB.Delete(&[]Program{{ID: 9}, {ID: 10}})
	// 若以其他字段作为查询条件
	// 参数conds以SQL语句的条件查询格式表示即可
	DB.Delete(&Program{}, "name = ?", "大闹天宫")
	DB.Where("name = ?", "姜子牙A").Delete(&Program{})
	// 查询已删除的数据
	p1 := Program{}
	DB.Unscoped().Where("id = ?", 6).Find(&p1)
	fmt.Printf("软删除的数据为：%v\n", p1.Name)
	// 永久删除数据
	DB.Unscoped().Delete(&p1)
}
