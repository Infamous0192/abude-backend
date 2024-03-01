package shift

import (
	"abude-backend/pkg/pagination"

	"gorm.io/datatypes"
)

type ShiftDTO struct {
	Name        string         `json:"name" form:"name" validate:"required"`
	Description string         `json:"description" form:"description" validate:"omitempty"`
	StartTime   datatypes.Time `json:"startTime" form:"startTime" validate:"required"`
	EndTime     datatypes.Time `json:"endTime" form:"endTime" validate:"required"`
	Company     uint           `json:"company" form:"company" validate:"required,exist=companies"`
}

type ShiftQuery struct {
	pagination.Pagination
	Keyword string `query:"keyword"`
	Company uint   `query:"company"`
}
