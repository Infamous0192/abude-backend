package product

import "abude-backend/pkg/pagination"

type IngredientDTO struct {
	Quantity float64 `json:"quantity" form:"quantity" validate:"required,min=0"`
	Product  uint    `json:"product" form:"product" validate:"required"`
}

type ProductDTO struct {
	Name        string  `json:"name" form:"name" validate:"required"`
	Description string  `json:"description" form:"description" validate:"omitempty"`
	Price       float64 `json:"price" form:"price" validate:"required"`
	Unit        string  `json:"unit" form:"unit" validate:"required"`
	Type        string  `json:"type" form:"type" validate:"required,oneof=purchase sale" enums:"purchase,sale"`
	Company     uint    `json:"company" form:"company" validate:"required,exist=companies"`
	Category    *uint   `json:"category" form:"category" validate:"omitempty,exist=categories"`
	IsDefault   bool    `json:"isDefault" form:"isDefault" validate:"omitempty"`
	Stock       bool    `json:"stock" form:"stock" validate:"required"`

	Ingredients []IngredientDTO `json:"ingredients" form:"ingredients" validate:"omitempty,dive,required"`
}

type ProductQuery struct {
	pagination.Pagination
	Keyword  string `query:"keyword"`
	Company  int    `query:"company"`
	Type     string `query:"type" enums:"purchase,sale"`
	Category int    `query:"category"`
	Default  *bool  `query:"default"`
}
