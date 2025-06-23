package models

import (
	"baby/settings"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, p int) (*gorm.DB, int, int, int, int) {
	// 当前页数小于或等于0，则当前页数变为第一页
	if p <= 0 {
		p = 1
	}
	// 计算所有数据总数和总页数
	var count int64
	db.Count(&count)
	pageCount := int(count) / settings.PageSize
	// 如果存在余数，则对pageCount加1处理
	if int(count)%settings.PageSize > 0 {
		pageCount += 1
	}
	// 当前页数超出总页数，则当前页数变为总页数
	if p >= pageCount {
		p = pageCount
	}
	// 计算上一页
	previous := 1
	if p >= 0 {
		previous = p - 1
	}
	// 计算下一页
	next := p + 1
	if next > pageCount {
		next = pageCount
	}
	// 计算数据偏移量，用于数据查询
	offset := (p - 1) * settings.PageSize
	res := db.Offset(offset).Limit(settings.PageSize)
	return res, previous, next, int(count), pageCount
}
