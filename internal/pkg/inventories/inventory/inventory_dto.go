package inventory

import (
	"abude-backend/pkg/pagination"
	"time"

	"gorm.io/datatypes"
)

type InventoryDTO struct {
	Source   string         `json:"source" form:"source" validate:"required,oneof=outlet warehouse"`
	SourceID uint           `json:"sourceId" form:"sourceId" validate:"required"`
	Date     datatypes.Date `json:"date" form:"date" validate:"required"`
	Product  uint           `json:"product" form:"product" validate:"required,exist=products.id"`
	Price    float64        `json:"price" form:"price" validate:"required,min=0"`
	Quantity float64        `json:"quantity" form:"quantity" validate:"required"`
}

type InventoryQuery struct {
	pagination.Pagination
	Product   int       `query:"product"`
	Outlet    int       `query:"outlet"`
	Warehouse int       `query:"warehouse"`
	StartDate time.Time `query:"startDate"`
	EndDate   time.Time `query:"endDate"`
}

type InventoryUpdateDTO struct {
	Date     datatypes.Date `json:"date" form:"date" validate:"required"`
	Quantity int64          `json:"quantity" form:"quantity" validate:"required"`
}

type StockQuery struct {
	pagination.Pagination
	Outlet  int `query:"outlet"`
	Product int `query:"product"`
}

type StockSummaryQuery struct {
	Outlet int `query:"outlet"`
}
