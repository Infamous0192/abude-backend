package expense

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/accounts/account"
	"abude-backend/internal/pkg/company"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/user"
	"abude-backend/pkg/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	TypeDebit  = "debit"
	TypeCredit = "credit"
)

const (
	StatusAccepted = "accepted"
	StatusApproved = "approved"
	StatusCanceled = "canceled"
)

type Expense struct {
	common.BaseModel
	user.WithEditor
	Code   string    `json:"code" gorm:"type:varchar(100)"`
	Amount float64   `json:"amount"`
	Type   string    `json:"type" gorm:"type:enum('debit','credit')" enums:"debit,credit"`
	Status string    `json:"status" gorm:"type:enum('accepted','approved','canceled')" enums:"approved,accepted,canceled"`
	Date   time.Time `json:"date"`
	Notes  string    `json:"notes"`

	Company   *company.Company `json:"company,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID uint             `json:"-"`

	Outlet   *outlet.Outlet `json:"outlet,omitempty" gorm:"constraint:OnDelete:SET NULL;"`
	OutletID uint           `json:"-"`

	Account   *account.Account `json:"account,omitempty" gorm:"constraint:OnDelete:RESTRICT;"`
	AccountID uint             `json:"-"`
}

func (expense *Expense) BeforeCreate(tx *gorm.DB) error {
	if expense.Code == "" {
		now := time.Now()

		var count int64
		tx.Model(&Expense{}).
			Where("DATE(created_at) = ?", now.Format("2006-01-02")).
			Count(&count)

		expense.Code = fmt.Sprintf("EXP-%s%s", now.Format("20060102"), utils.NumberToDigit(int(count+1), 3))
	}

	return nil
}

type ExpenseSummary struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Quantity float64 `json:"quantity"`
	Total    float64 `json:"total"`
}
