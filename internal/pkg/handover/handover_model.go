package handover

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/attendances/shift"
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/transactions/expense"
	"abude-backend/internal/pkg/transactions/purchase"
	"abude-backend/internal/pkg/transactions/sale"
	"abude-backend/internal/pkg/user"
	"time"
)

type HandoverItem struct {
	Price     int64           `json:"price"`
	Quantity  int64           `json:"quantity"`
	Total     int64           `json:"total"`
	Product   product.Product `json:"product"`
	ProductID uint            `json:"-"`
	Date      time.Time       `json:"date"`
}

type Handover struct {
	common.BaseModel
	user.WithEditor
	Note           string    `json:"note" gorm:"type:varchar(150)"`
	SalesTotal     float64   `json:"salesTotal"`
	PurchasesTotal float64   `json:"purchasesTotal"`
	ExpensesTotal  float64   `json:"expensesTotal"`
	CashReceived   float64   `json:"cashReceived"`
	CashReturned   float64   `json:"cashReturned"`
	Date           time.Time `json:"date"`

	Sales     []sale.Sale         `json:"-" gorm:"many2many:handover_sales;constraint:OnDelete:CASCADE;"`
	Purchases []purchase.Purchase `json:"-" gorm:"many2many:handover_purchases;constraint:OnDelete:CASCADE;"`
	Expenses  []expense.Expense   `json:"-" gorm:"many2many:handover_expenses;constraint:OnDelete:CASCADE;"`

	SaleItems     []HandoverItem `json:"sales" gorm:"-"`
	PurchaseItems []HandoverItem `json:"purchases" gorm:"-"`
	ExpenseItems  []HandoverItem `json:"expenses" gorm:"-"`

	Outlet   *outlet.Outlet `json:"outlet,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint           `json:"-"`

	Shift   *shift.Shift `json:"shift,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	ShiftID uint         `json:"-"`
}
