package sale

import (
	"abude-backend/pkg/pagination"
	"time"
)

type SaleItemDTO struct {
	Price    *float64 `json:"price" form:"price" validate:"omitempty"`
	Quantity float64  `json:"quantity" form:"quantity" validate:"required"`
	Product  uint     `json:"product" form:"product" validate:"required,exist=products"`
}

type SaleDTO struct {
	Customer string        `json:"customer" form:"customer" validate:"required"`
	Note     string        `json:"note" form:"note" validate:"omitempty"`
	Items    []SaleItemDTO `json:"items" form:"items" validate:"required,dive,required"`
	Source   string        `json:"source" form:"source" validate:"required,oneof=outlet warehouse"`
	SourceID uint          `json:"sourceId" form:"sourceId" validate:"required"`
	Date     time.Time     `json:"date" form:"date" validate:"omitempty" format:"date-time"`
	Status   *string       `json:"status" form:"status" validate:"omitempty,oneof=accepted canceled approved" enums:"accepted,canceled,approved"`

	User uint `json:"-" form:"-"`
}

type SaleQuery struct {
	pagination.Pagination
	User      string    `query:"user"`   // User ID
	Outlet    uint      `query:"outlet"` // Outlet ID
	Status    []string  `query:"status" enums:"accepted,approved,canceled"`
	StartDate time.Time `query:"startDate" format:"date-time"`
	EndDate   time.Time `query:"endDate" format:"date-time"`
}

type SaleSummaryQuery struct {
	Status    []string `query:"status" enums:"accepted,approved,canceled"`
	Outlet    uint     `query:"outlet"` // Outlet ID
	StartDate string   `query:"startDate" format:"date-time"`
	EndDate   string   `query:"endDate" format:"date-time"`
}
