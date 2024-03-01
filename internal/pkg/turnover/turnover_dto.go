package turnover

import (
	"abude-backend/pkg/pagination"
	"time"
)

type TurnoverDTO struct {
	Date     time.Time `form:"date" json:"date" validate:"required"`
	Evidence string    `form:"evidence" json:"evidence" validate:"required,url"`
	Outlet   uint      `form:"outlet" json:"outlet" validate:"required,exist=outlets"`

	User *uint `form:"-" json:"-"`
}

type TurnoverQuery struct {
	pagination.Pagination
	Outlet    uint      `query:"outlet"`
	StartDate time.Time `query:"startDate"`
	EndDate   time.Time `query:"endDate"`
}
