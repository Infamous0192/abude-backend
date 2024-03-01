package company

import "abude-backend/pkg/pagination"

type CompanyDTO struct {
	Name   string `form:"name" json:"name" validate:"required"`
	Region string `form:"region" json:"region" validate:"required"`
}

type CompanyQuery struct {
	pagination.Pagination
	Keyword string `query:"keyword"`
	Owner   []uint `query:"owner"`
}
