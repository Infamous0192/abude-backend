package supplier

import "abude-backend/pkg/pagination"

type SupplierDTO struct {
	Name        string `form:"name" json:"name" validate:"required"`
	Description string `form:"description" json:"description" validate:"omitempty"`
	Company     uint   `form:"company" json:"company" validate:"required,exist=companies"`
}

type SupplierQuery struct {
	pagination.Pagination
	Keyword string `query:"keyword"`
	Company uint   `query:"company"`
}
