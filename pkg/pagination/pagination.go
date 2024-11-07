package pagination

import (
	"gorm.io/gorm"
	"math"
)

// 	var lists []entity.ChargeDiscount
//	db := master.
//		Model(entity.ChargeDiscount{}).
//		Where(where).
//		Where("status < ?", entity.ChargeDiscountStatusOff).
//		Order("ID DESC")
//
//	pager := pagination.Page(db, pageNo, pagination.WithPageSize(pageSize))
//	pager.DB.Find(&lists)
//	pager.List = lists
//	return pager

// 分页相关
const (
	PageSize = 20 // 每页大小
)

// PageParams 分页参数
type PageParams struct {
	Page     int `json:"pagination" form:"pagination"` // 请求页数
	PageSize int `json:"page_size" form:"page_size"`   // 每页大小,最大 20, 大于 20 以 20 计算
}

// Op 分页配置项
type Op struct {
	PageSize int // 每页数量
}

// Option 参数类型
type Option func(*Op)

// applyOpts 应用可选参数
func (op *Op) applyOpts(opts []Option) {
	for _, opt := range opts {
		opt(op)
	}
}

// WithPageSize 每页数量
func WithPageSize(size int) Option {
	return func(op *Op) { op.PageSize = size }
}

// Next 获取下一页页数
func Next(total int64, pageSize, currPage int) (nextPage, totalPage int) {
	if total <= 0 {
		return
	}
	if pageSize <= 0 || currPage <= 0 {
		return 0, 1
	}
	totalPageCount := math.Ceil(float64(total) / float64(pageSize))
	if totalPageCount-float64(currPage) > 0 {
		nextPage = currPage + 1
	}
	return nextPage, int(totalPageCount)
}

type Pager struct {
	DB    *gorm.DB    `json:"-"`
	Pages int         `json:"pages"` // 总页数
	Next  int         `json:"next"`  // 下一页;没有则为0
	Cur   int         `json:"cur"`   //当前页码
	Count int64       `json:"count"` // 总条数
	List  interface{} `json:"list"`  // 列表数据
	Limit int         `json:"limit"` // 每页条数
}

// Page 分页
// @param db *gorm.db
// @param pageNo PageParams.Page
func Page(db *gorm.DB, pageNo int, opts ...Option) (pager Pager) {
	if pageNo <= 0 {
		pageNo = 1
	}

	op := Op{}
	op.applyOpts(opts)
	if op.PageSize == 0 || op.PageSize > 6000 { //每页条数控制
		op.PageSize = PageSize
	}

	var count int64
	db.Count(&count)
	if count == 0 {
		return Pager{
			DB:    db,
			Pages: 0,
			Next:  0,
			Cur:   0,
			Count: 0,
		}
	}

	next, pages := Next(count, op.PageSize, pageNo)
	if op.PageSize > 0 && pageNo > 0 {
		db = db.Offset(op.PageSize * (pageNo - 1)).Limit(op.PageSize)
	}
	return Pager{
		DB:    db,
		Pages: pages,
		Next:  next,
		Cur:   pageNo,
		Count: count,
	}
}
