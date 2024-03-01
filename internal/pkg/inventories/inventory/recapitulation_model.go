package inventory

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type RecapitulationItem struct {
	common.BaseModel
	Available  float64 `json:"available"`
	TotalValue float64 `json:"totalValue"`
	StockIn    float64 `json:"stockIn"`
	ValueIn    float64 `json:"valueIn"`
	StockOut   float64 `json:"stockOut"`
	ValueOut   float64 `json:"valueOut"`

	Product   *product.Product `json:"product" gorm:"constraint:OnDelete:RESTRICT;"`
	ProductID uint             `json:"-"`

	Recapitulation   Recapitulation `json:"recapitulation" gorm:"constraint:OnDelete:RESTRICT;"`
	RecapitulationID uint           `json:"-"`
}

func (RecapitulationItem) TableName() string {
	return "inventory_recap_items"
}

type Recapitulation struct {
	common.BaseModel
	user.WithEditor
	Code     string    `json:"code" gorm:"type:varchar(100)"`
	Date     time.Time `json:"date"`
	Notes    string    `json:"notes"`
	Employee string    `json:"employee"`

	Items []RecapitulationItem `json:"items"`

	Outlet   *outlet.Outlet `json:"outlet" gorm:"constraint:OnDelete:RESTRICT;"`
	OutletID uint           `json:"-"`
}

func (Recapitulation) TableName() string {
	return "inventory_recaps"
}

func (recap *Recapitulation) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	var count int64
	tx.Model(&Recapitulation{}).
		Where("DATE(created_at) = ?", now.Format("2006-01-02")).
		Count(&count)

	recap.Code = fmt.Sprintf("STX-%s%d%s", now.Format("20060102"), recap.OutletID, utils.NumberToDigit(int(count+1), 3))

	return nil
}
