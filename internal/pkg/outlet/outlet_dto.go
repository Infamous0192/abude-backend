package outlet

import "abude-backend/pkg/pagination"

type OutletDTO struct {
	Name      string  `json:"name" form:"name" validate:"required"`
	Address   string  `json:"address" form:"address" validate:"omitempty"`
	Latitude  float64 `json:"latitude" form:"latitude" validate:"omitempty,latitude"`
	Longitude float64 `json:"longitude" form:"longitude" validate:"omitempty,longitude"`
	Status    *bool   `json:"status" form:"status" validate:"required"`
	Company   uint    `json:"company" form:"company" validate:"required,exist=companies"`
}

type OutletQuery struct {
	pagination.Pagination
	Status   *bool  `query:"status"`
	Keyword  string `query:"keyword"`
	Owner    uint   `query:"owner"`
	Company  uint   `query:"company"`
	Employee uint   `query:"admin"`
}

type OutletCountQuery struct {
	Company uint `query:"company"`
}

type OutletEmployeeDTO struct {
	Type     string `json:"type" form:"type" validate:"required,oneof=admin employee"`
	Employee int    `json:"-" form:"-" validate:"required,exist=employees"`
	Outlet   int    `json:"-" form:"-" validate:"required,exist=outlets"`
}

type OutletEmployeeQuery struct {
	pagination.Pagination
	Type     string `query:"type" enums:"admin,employee"`
	Employee int    `query:"employee"`
	Outlet   int    `query:"-"`
}
