package company

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/user"
)

type Company struct {
	common.BaseModel
	Name   string `json:"name" gorm:"type:varchar(100)"`
	Region string `json:"region" gorm:"type:varchar(100)"`

	Owners []user.User `json:"-" gorm:"many2many:company_owners;constraint:OnDelete:CASCADE;"`
}
