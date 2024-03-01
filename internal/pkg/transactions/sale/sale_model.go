package sale

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

const (
	StatusAccepted = "accepted"
	StatusApproved = "approved"
	StatusCanceled = "canceled"
)

type SaleItem struct {
	common.BaseModel
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Total    float64 `json:"total"`
	Status   bool    `json:"status"`

	Product   *product.Product `json:"product,omitempty" gorm:"constraint:OnDelete:RESTRICT;"`
	ProductID uint             `json:"-"`

	Sale   *Sale `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	SaleID uint  `json:"-"`
}

type Sale struct {
	common.BaseModel
	Code     string    `json:"code" gorm:"type:varchar(50)"`
	Note     string    `json:"note" gorm:"type:varchar(150)"`
	Customer string    `json:"customer"`
	Total    float64   `json:"total"`
	Status   string    `json:"status" gorm:"type:enum('accepted','approved','canceled')" enums:"approved,accepted,canceled"`
	Date     time.Time `json:"date"`

	Items []SaleItem `json:"items" gorm:"constraint:OnDelete:CASCADE;"`

	User   *user.User `json:"user,omitempty"`
	UserID uint       `json:"-"`
}

type SaleSummary struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Quantity float64 `json:"quantity"`
	Total    float64 `json:"total"`
}

type SaleTest struct {
	Product product.Product `json:"product" gorm:"embedded"`
	Total   float64         `json:"total"`
}

type OutletSale struct {
	Outlet   *outlet.Outlet `gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint           `gorm:"primaryKey"`

	Sale   *Sale `gorm:"constraint:OnDelete:CASCADE;"`
	SaleID uint  `gorm:"primaryKey"`
}

func (sale *Sale) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()

	var count int64
	tx.Model(&Sale{}).
		Where("DATE(created_at) = ?", now.Format("2006-01-02")).
		Count(&count)

	sale.Code = fmt.Sprintf("TXS-%s%s", now.Format("20060102"), utils.NumberToDigit(int(count+1), 3))

	return nil
}
