package common

import (
	"abude-backend/pkg/validation"

	"gorm.io/gorm"
)

type BaseController struct {
	Validation *validation.Validation
	DB         *gorm.DB
}
