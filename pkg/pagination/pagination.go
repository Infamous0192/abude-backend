package pagination

import (
	"abude-backend/pkg/exception"
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Page  int `query:"page" json:"page"`
	Limit int `query:"limit" json:"limit"`
}

type Metadata struct {
	Pagination
	Total   int64 `json:"total"`
	Count   int   `json:"count"`
	HasPrev bool  `json:"hasPrev"`
	HasNext bool  `json:"hasNext"`
}

type Result[T interface{}] struct {
	Metadata Metadata `json:"metadata"`
	Result   []T      `json:"result"`
}

func New[T interface{}](pagination Pagination) Result[T] {
	return Result[T]{
		Metadata: Metadata{
			Pagination: pagination,
		},
	}
}

func (result *Result[T]) Paginate(db *gorm.DB) *Result[T] {
	metadata := &result.Metadata

	if metadata.Page == 0 {
		metadata.Page = 1
	}

	if metadata.Limit == 0 {
		metadata.Limit = 5
	}

	tx := db.Session(&gorm.Session{})
	tx.Count(&metadata.Total)

	offset := metadata.Limit * (metadata.Page - 1)
	err := db.Offset(offset).Limit(metadata.Limit).Find(&result.Result).Error
	exception.CatchDB(err)

	metadata.Count = len(result.Result)
	metadata.HasPrev = metadata.Page > 1
	metadata.HasNext = metadata.Count > 0 && metadata.Page < int(math.Ceil(float64(metadata.Total)/float64(metadata.Limit)))

	return result
}
