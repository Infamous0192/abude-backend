package employee

import "abude-backend/pkg/pagination"

type EmployeeCreateDTO struct {
	Username    string `form:"username" json:"username" validate:"required"`
	Password    string `form:"password" json:"password" validate:"required,min=6"`
	Name        string `form:"name" json:"name" validate:"required"`
	Phonenumber string `form:"phonenumber" json:"phonenumber" validate:"omitempty,e164"`
	Address     string `form:"address" json:"address" validate:"omitempty"`
	Status      *bool  `form:"status" json:"status" validate:"required"`
}

type EmployeeUpdateDTO struct {
	Username    string `form:"username" json:"username" validate:"required"`
	Password    string `form:"password" json:"password" validate:"omitempty,min=6"`
	Name        string `form:"name" json:"name" validate:"required"`
	Phonenumber string `form:"phonenumber" json:"phonenumber" validate:"omitempty,e164"`
	Address     string `form:"address" json:"address" validate:"omitempty"`
	Status      *bool  `form:"status" json:"status" validate:"required"`
}

type EmployeeQuery struct {
	pagination.Pagination
	Keyword string `query:"keyword"`
	Outlet  []uint `query:"outlet"`
}
