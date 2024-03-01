package handover

import (
	"abude-backend/pkg/pagination"
	"time"
)

type ProofDTO struct {
	Date     time.Time `form:"date" json:"date" validate:"required"`
	Evidence string    `form:"evidence" json:"evidence" validate:"required,url"`
	Outlet   uint      `form:"outlet" json:"outlet" validate:"required,exist=outlets"`
}

type ProofQuery struct {
	pagination.Pagination
	Outlet    uint      `query:"outlet"`
	StartDate time.Time `query:"startDate"`
	EndDate   time.Time `query:"endDate"`
}
