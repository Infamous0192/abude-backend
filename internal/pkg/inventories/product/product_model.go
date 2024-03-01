package product

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/company"
	"abude-backend/internal/pkg/inventories/category"
)

type Ingredient struct {
	common.BaseModel
	Quantity float64 `json:"quantity"`

	Base   *Product `json:"base" gorm:"constraint:OnDelete:CASCADE;"`
	BaseID uint     `json:"-"`

	Ingredient   *Product `json:"ingredient" gorm:"constraint:OnDelete:RESTRICT;"`
	IngredientID uint     `json:"-"`
}

type Product struct {
	common.BaseModel
	Name        string  `json:"name" gorm:"type:varchar(100)"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Unit        string  `json:"unit" gorm:"type:varchar(100)"`
	Type        string  `json:"type" gorm:"type:enum('purchase','sale')" enums:"purchase,sale"`
	IsDefault   bool    `json:"isDefault"`
	Stock       bool    `json:"stock"`

	Ingredients []Ingredient `json:"ingredients" gorm:"foreignKey:base_id"`

	Category   *category.Category `json:"category" gorm:"constraint:OnDelete:SET NULL;"`
	CategoryID *uint              `json:"-"`

	Company   *company.Company `json:"company" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID uint             `json:"-"`
}
