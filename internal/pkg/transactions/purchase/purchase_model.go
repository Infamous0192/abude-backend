package purchase

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/supplier"
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	StatusAccepted = "accepted"
	StatusApproved = "approved"
	StatusCanceled = "canceled"
)

const (
	TypeDebit  = "debit"
	TypeCredit = "credit"
)

type PurchaseItem struct {
	common.BaseModel
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Total    float64 `json:"total"`
	Status   bool    `json:"status"`

	Product   *product.Product `json:"product,omitempty" gorm:"constraint:OnDelete:RESTRICT;"`
	ProductID uint             `json:"-"`

	Purchase   Purchase `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	PurchaseID uint     `json:"-"`
}

type Purchase struct {
	common.BaseModel
	Code   string    `json:"code" gorm:"type:varchar(50)"`
	Note   string    `json:"note" gorm:"type:varchar(150)"`
	Total  float64   `json:"total"`
	Status string    `json:"status" gorm:"type:enum('accepted','approved','canceled')" enums:"approved,accepted,canceled"`
	Type   string    `json:"type" gorm:"type:enum('debit','credit')" enums:"debit,credit"`
	Date   time.Time `json:"date"`

	Items []PurchaseItem `json:"items" gorm:"constraint:OnDelete:CASCADE;"`

	Supplier   *supplier.Supplier `json:"supplier,omitempty" gorm:"constraint:OnDelete:RESTRICT;"`
	SupplierID *uint              `json:"-"`

	User   *user.User `json:"user,omitempty" gorm:"constraint:OnDelete:SET NULL;"`
	UserID uint       `json:"-"`
}

type PurchaseSummary struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Quantity float64 `json:"quantity"`
	Total    float64 `json:"total"`
}

type OutletPurchase struct {
	Outlet   *outlet.Outlet `gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint           `gorm:"primaryKey"`

	Purchase   *Purchase `gorm:"constraint:OnDelete:CASCADE;"`
	PurchaseID uint      `gorm:"primaryKey"`
}

func (purchase *Purchase) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	var count int64
	err := tx.Model(&Purchase{}).
		Where("DATE(created_at) = ?", now.Format("2006-01-02")).
		Count(&count).Error

	if err != nil {
		return err
	}

	purchase.Code = fmt.Sprintf("TXP-%s%s", now.Format("20060102"), utils.NumberToDigit(int(count+1), 3))

	return nil
}
