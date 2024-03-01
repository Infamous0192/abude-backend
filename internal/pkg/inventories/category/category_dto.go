package category

import "abude-backend/pkg/pagination"

type CategoryDTO struct {
	Name        string `form:"name" json:"name" validate:"required"`
	Description string `form:"description" json:"description" validate:"omitempty"`
	Parent      *uint  `form:"parent" json:"parent" validate:"omitempty,exist=categories"`
}

type CategoryQuery struct {
	pagination.Pagination
	Parent int `query:"parent"`
}
