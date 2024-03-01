package wage

import (
	"abude-backend/pkg/pagination"
	"time"
)

type WageDTO struct {
	Amount   float64   `json:"amount" form:"amount" validate:"required,min=0"`
	Type     string    `json:"type" form:"type" validate:"required,oneof=debit credit" enums:"debit,credit"`
	Date     time.Time `json:"date" form:"date" validate:"omitempty"`
	Notes    string    `json:"notes" form:"notes" validate:"omitempty"`
	Employee uint      `json:"employee" form:"employee" validate:"required,exist=employees"`
	Company  uint      `json:"company" form:"company" validate:"required,exist=companies"`
	Outlet   uint      `json:"outlet" form:"outlet" validate:"required,exist=outlets"`
}

type WageQuery struct {
	pagination.Pagination
	Employee uint `query:"employee"`
	Company  uint `query:"company"`
	Outlet   uint `query:"outlet"`
}

type WageSummaryQuery struct {
	Status    []string `query:"status" enums:"accepted,approved,canceled"`
	Outlet    uint     `query:"outlet"`  // Outlet ID
	Company   uint     `query:"company"` // Company ID
	Employee  uint     `query:"employee"`
	StartDate string   `query:"startDate" format:"date-time"`
	EndDate   string   `query:"endDate" format:"date-time"`
}
