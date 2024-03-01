package account

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/accounts/category"
	"abude-backend/internal/pkg/company"
	"abude-backend/pkg/exception"

	"gorm.io/gorm"
)

type Account struct {
	common.BaseModel
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description"`
	Code        string `json:"code" gorm:"type:varchar(6)"`
	Normal      int    `json:"normal" gorm:"type:tinyint(1)"`

	Category   category.Category `json:"category" gorm:"constraint:OnDelete:RESTRICT;"`
	CategoryID uint              `json:"-"`

	Company   *company.Company `json:"company,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID *uint            `json:"-"`
}

func (account *Account) BeforeSave(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Account{}).
		Where("code = ? AND company_id = ? AND id != ?", account.Code, account.CompanyID, account.ID).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return exception.Validation(map[string]string{
			"code": "Kode telah digunakan",
		})
	}

	return nil
}
