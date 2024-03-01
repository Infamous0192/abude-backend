package expense

import (
	"abude-backend/pkg/pagination"
	"time"
)

type ExpenseDTO struct {
	Amount  float64   `json:"amount" form:"amount" validate:"required,min=0"`
	Type    string    `json:"type" form:"type" validate:"required,oneof=debit credit" enums:"debit,credit"`
	Date    time.Time `json:"date" form:"date" validate:"omitempty"`
	Notes   string    `json:"notes" form:"notes" validate:"omitempty"`
	Account uint      `json:"account" form:"account" validate:"required,exist=accounts"`
	Company uint      `json:"company" form:"company" validate:"required,exist=companies"`
	Outlet  uint      `json:"outlet" form:"outlet" validate:"required,exist=outlets"`
}

type ExpenseQuery struct {
	pagination.Pagination
	Status  []string `query:"status"`
	Account uint     `query:"account"`
	Company uint     `query:"company"`
	Outlet  uint     `query:"outlet"`
}

type ExpenseSummaryQuery struct {
	Status    []string `query:"status" enums:"accepted,approved,canceled"`
	Outlet    uint     `query:"outlet"`  // Outlet ID
	Company   uint     `query:"company"` // Company ID
	Account   uint     `query:"account"`
	StartDate string   `query:"startDate" format:"date-time"`
	EndDate   string   `query:"endDate" format:"date-time"`
}
