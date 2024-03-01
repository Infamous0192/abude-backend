package user

import "abude-backend/pkg/pagination"

type UserCreateDTO struct {
	Name     string `form:"name" json:"name" validate:"required"`
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required,min=6"`
	Role     string `form:"role" json:"role" validate:"required,oneof=superadmin owner employee"`
	Status   *bool  `form:"status" json:"status" validate:"required"`
}

type UserUpdateDTO struct {
	Name     string `form:"name" json:"name" validate:"required"`
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"omitempty,min=6"`
	Role     string `form:"role" json:"role" validate:"required,oneof=superadmin owner employee"`
	Status   *bool  `form:"status" json:"status" validate:"required"`
}

type UserQuery struct {
	pagination.Pagination
	Role []string `query:"role" enums:"superadmin,owner,employee"`
}
