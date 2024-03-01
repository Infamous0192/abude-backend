package turnover

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/user"
	"time"

	"gorm.io/gorm"
)

type Turnover struct {
	common.BaseModel
	user.WithEditor
	Date     time.Time `json:"date"`
	Evidence string    `json:"evidence"`
	Income   int64     `json:"income" gorm:"-"`
	Expense  int64     `json:"expense" gorm:"-"`

	Outlet   *outlet.Outlet `json:"outlet,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint           `json:"-"`
}

func (t *Turnover) AfterFind(tx *gorm.DB) error {
	var income int64
	tx.Table("sales").
		Select("SUM(total)").
		Joins("INNER JOIN outlet_sales ON sales.id=outlet_sales.sale_id").
		Where("outlet_sales.outlet_id = ? AND DATE(sales.date) = DATE(?) AND sales.status != 'canceled'", t.OutletID, t.Date).Row().Scan(&income)

	t.Income = income

	var expense int64
	tx.Table("purchases").
		Select("SUM(total)").
		Joins("INNER JOIN outlet_purchases ON purchases.id=outlet_purchases.purchase_id").
		Where("outlet_purchases.outlet_id = ? AND DATE(purchases.date) = DATE(?) AND purchases.status != 'canceled'", t.OutletID, t.Date).Row().Scan(&expense)

	t.Expense = expense

	return nil
}
