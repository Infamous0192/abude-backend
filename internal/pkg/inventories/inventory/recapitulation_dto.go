package inventory

import (
	"abude-backend/pkg/pagination"
	"time"
)

type RecapitulationDTO struct {
	Notes    string    `json:"notes" form:"notes" validate:"omitempty"`
	Employee string    `json:"employee" form:"employee" validate:"required"`
	Date     time.Time `json:"date" form:"date" validate:"required"`
	Outlet   uint      `json:"outlet" form:"outlet" validate:"required,exist=outlets"`
}

type RecapitulationQuery struct {
	pagination.Pagination
	Keyword string `query:"keyword"`
	Outlet  int    `query:"outlet"`
}
