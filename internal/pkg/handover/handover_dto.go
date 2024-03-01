package handover

import (
	"abude-backend/pkg/pagination"
	"time"
)

type HandoverDTO struct {
	Note         string    `form:"note" json:"note" validate:"required"`
	CashReceived float64   `form:"cashReceived" json:"cashReceived" validate:"required"`
	CashReturned float64   `form:"cashReturned" json:"cashReturned" validate:"omitempty"`
	Outlet       uint      `form:"outlet" json:"outlet" validate:"required,exist=outlets"`
	Shift        uint      `form:"shift" json:"shift" validate:"required,exist=shifts"`
	Date         time.Time `form:"date" json:"date" validate:"required" format:"date-time"`
}

type HandoverQuery struct {
	pagination.Pagination
	Outlet    uint      `query:"outlet"`
	Shift     uint      `query:"shift"`
	StartDate time.Time `query:"startDate" format:"date-time"`
	EndDate   time.Time `query:"endDate" format:"date-time"`
}
