package pkg

import "gorm.io/gorm"

type paginate struct {
	limit int
	page  int
}

func NewPaginate(limit int, page int) *paginate {
	return &paginate{limit: limit, page: page}
}

func (p *paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Offset(offset).
		Limit(p.limit)
}

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/jinzhu/gorm"
// )

// func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		q := r.URL.Query()
// 		page, _ := strconv.Atoi(q.Get("page"))
// 		if page <= 0 {
// 			page = 1
// 		}

// 		pageSize, _ := strconv.Atoi(q.Get("page_size"))
// 		switch {
// 		case pageSize > 100:
// 			pageSize = 100
// 		case pageSize <= 0:
// 			pageSize = 10
// 		}

// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }

// import (
// 	"math"

// 	"github.com/jinzhu/gorm"
// )

// type Pagination struct {
// 	Limit      int         `json:"limit,omitempty;query:limit"`
// 	Page       int         `json:"page,omitempty;query:page"`
// 	Sort       string      `json:"sort,omitempty;query:sort"`
// 	TotalRows  int64       `json:"total_rows"`
// 	TotalPages int         `json:"total_pages"`
// 	Rows       interface{} `json:"rows"`
// }

// func (p *Pagination) GetOffset() int {
// 	return (p.GetPage() - 1) * p.GetLimit()
// }

// func (p *Pagination) GetLimit() int {
// 	if p.Limit == 0 {
// 		p.Limit = 10
// 	}
// 	return p.Limit
// }

// func (p *Pagination) GetPage() int {
// 	if p.Page == 0 {
// 		p.Page = 1
// 	}
// 	return p.Page
// }

// func (p *Pagination) GetSort() string {
// 	if p.Sort == "" {
// 		p.Sort = "Id desc"
// 	}
// 	return p.Sort
// }

// func paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
// 	var totalRows int64
// 	db.Model(value).Count(&totalRows)

// 	pagination.TotalRows = totalRows
// 	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
// 	pagination.TotalPages = totalPages

// 	return func(db *gorm.DB) *gorm.DB {
// 		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
// 	}
// }
