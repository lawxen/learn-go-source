package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Jwts struct {
	gorm.Model
	Token  string    `json:"token" gorm:"type:varchar(1000)"`
	Expire time.Time `json:"expire"`
}

// 定义数据库连接对象
var dsn = "root:1234@tcp(127.0.0.1:3306)/baby?charset=utf8mb4&parseTime=True&loc=Local"
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 禁止创建数据表的外键约束
	DisableForeignKeyConstraintWhenMigrating: true,
})

// 创建Casbin对象
var A, _ = gormadapter.NewAdapterByDB(DB)
var E, _ = casbin.NewEnforcer("middleware/model.conf", A)

// Setup initializes the database instance
func Setup() {
	if err != nil {
		fmt.Printf("模型初始化异常: %v", err)
	}
	// 数据迁移
	DB.AutoMigrate(&Jwts{})
	// 加载策略
	E.LoadPolicy()
	// 设置数据库连接池
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}
