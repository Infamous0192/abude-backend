package employee

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/user"
)

type Employee struct {
	common.BaseModel
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Phonenumber string `json:"phonenumber" gorm:"type:varchar(25)"`
	Address     string `json:"address" gorm:"type:varchar(150)"`
	Status      bool   `json:"status"`

	User   *user.User `json:"user,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	UserID uint       `json:"-"`
}
