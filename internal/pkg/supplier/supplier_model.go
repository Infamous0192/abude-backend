package supplier

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/company"
)

type Supplier struct {
	common.BaseModel
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:varchar(255)"`

	Company   *company.Company `json:"company,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID uint             `json:"-"`
}
