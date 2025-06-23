package models

import (
	"baby/settings"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 定义数据表结构
type Types struct {
	gorm.Model
	Firsts  string `json:"firsts" gorm:"type:varchar(255)"`
	Seconds string `json:"seconds" gorm:"type:varchar(255)"`
}

type Commodities struct {
	gorm.Model
	Name     string    `json:"name" gorm:"type:varchar(255)"`
	Size     string    `json:"size" gorm:"type:varchar(255)"`
	Types    string    `json:"types" gorm:"type:varchar(255)"`
	Price    float64   `json:"price"`
	Discount float64   `json:"discount"`
	Stock    int64     `json:"stock"`
	Sold     int64     `json:"sold"`
	Likes    int64     `json:"likes"`
	Created  time.Time `json:"created"`
	Img      string    `json:"img" gorm:"type:varchar(255)"`
	Details  string    `json:"details" gorm:"type:varchar(255)"`
}

type Users struct {
	gorm.Model
	Username  string    `json:"username" gorm:"type:varchar(255);unique"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	IsStaff   int64     `json:"isStaff" gorm:"default:0"`
	LastLogin time.Time `json:"lastLogin"`
}

type Carts struct {
	gorm.Model
	Quantity    int64       `json:"quantity"`
	CommodityId int64       `json:"commodityId"`
	Commodities Commodities `gorm:"foreignkey:CommodityId"`
	UserId      int64       `json:"userId"`
	Users       Users       `json:"-" gorm:"foreignkey:UserId"`
}

type Orders struct {
	gorm.Model
	Price   string `json:"price" gorm:"type:varchar(255)"`
	PayInfo string `json:"payInfo" gorm:"type:varchar(255)"`
	UserId  int64
	Users   Users `json:"-" gorm:"foreignkey:UserId"`
	State   int64 `json:"state"`
}

type Records struct {
	gorm.Model
	CommodityId int64       `json:"commodityId"`
	Commodities Commodities `gorm:"foreignkey:CommodityId"`
	UserId      int64       `json:"userId"`
	Users       Users       `json:"-" gorm:"foreignkey:UserId"`
}

type Jwts struct {
	gorm.Model
	Token  string    `json:"token" gorm:"type:varchar(1000)"`
	Expire time.Time `json:"expire"`
}

// 定义模型Users的钩子函数BeforeCreate，为字段Password加密处理
func (u *Users) BeforeSave(db *gorm.DB) error {
	m := md5.New()
	m.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(m.Sum(nil))
	return nil
}

// 定义数据库连接对象
var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	settings.MySQLSetting.User,
	settings.MySQLSetting.Password,
	settings.MySQLSetting.Host,
	settings.MySQLSetting.Name)
var DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 禁止创建数据表的外键约束
	DisableForeignKeyConstraintWhenMigrating: true,
})

// Setup initializes the database instance
func Setup() {
	if err != nil {
		fmt.Printf("模型初始化异常: %v", err)
	}
	// 数据迁移
	DB.AutoMigrate(&Types{})
	DB.AutoMigrate(&Commodities{})
	DB.AutoMigrate(&Users{})
	DB.AutoMigrate(&Carts{})
	DB.AutoMigrate(&Orders{})
	DB.AutoMigrate(&Records{})
	DB.AutoMigrate(&Jwts{})
	// 设置数据库连接池
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
}
