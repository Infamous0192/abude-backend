package shift

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/company"

	"gorm.io/datatypes"
)

type Shift struct {
	common.BaseModel
	Name        string         `json:"name"`
	Description string         `json:"description"`
	StartTime   datatypes.Time `json:"startTime" gorm:"type:time"`
	EndTime     datatypes.Time `json:"endTime" gorm:"type:time"`

	Company   *company.Company `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID uint             `json:"-"`
}
