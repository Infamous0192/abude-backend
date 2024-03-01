package category

import "abude-backend/internal/common"

type Category struct {
	common.BaseModel
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description"`

	Parent   *Category `json:"parent" gorm:"constraint:OnDelete:CASCADE;"`
	ParentID *uint     `json:"-"`
}
