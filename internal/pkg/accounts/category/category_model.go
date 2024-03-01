package category

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/company"
	"abude-backend/pkg/exception"

	"gorm.io/gorm"
)

type Category struct {
	common.BaseModel
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description"`
	Code        string `json:"code" gorm:"type:varchar(2)"`
	Normal      int    `json:"normal" gorm:"type:tinyint(1)"`

	Company   *company.Company `json:"company,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID *uint            `json:"-"`
}

func (Category) TableName() string {
	return "account_categories"
}

func (category *Category) BeforeSave(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Category{}).
		Where("code = ? AND company_id = ? AND id != ?", category.Code, category.CompanyID, category.ID).
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
