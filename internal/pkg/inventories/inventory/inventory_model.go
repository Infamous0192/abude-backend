package inventory

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/inventories/category"
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"

	"gorm.io/datatypes"
)

type Inventory struct {
	common.BaseModel
	Date     datatypes.Date `json:"date"`
	StockIn  float64        `json:"stockIn"`
	StockOut float64        `json:"stockOut"`
	Price    float64        `json:"price"`

	Product   *product.Product `json:"product" gorm:"constraint:OnDelete:RESTRICT;"`
	ProductID uint             `json:"-"`
}

type OutletInventory struct {
	Inventory   *Inventory `json:"inventory" gorm:"constraint:OnDelete:CASCADE;"`
	InventoryID uint       `json:"-"`

	Outlet   *outlet.Outlet `json:"outlet" gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint           `json:"-"`
}

type Stock struct {
	Product      product.Product   `json:"product" gorm:"embedded"`
	Category     category.Category `json:"category" gorm:"embedded"`
	Amount       float64           `json:"amount"`
	TotalValue   float64           `json:"totalValue"`
	AveragePrice float64           `json:"averagePrice"`
}

type StockSummary struct {
	Product    product.Product `json:"product" gorm:"embedded"`
	Available  float64         `json:"available"`
	TotalValue float64         `json:"totalValue"`
	StockIn    float64         `json:"stockIn" gorm:"column:stock_in"`
	ValueIn    float64         `json:"valueIn" gorm:"column:value_in"`
	StockOut   float64         `json:"stockOut" gorm:"column:stock_out"`
	ValueOut   float64         `json:"valueOut" gorm:"column:value_out"`
}
