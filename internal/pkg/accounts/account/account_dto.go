package account

import "abude-backend/pkg/pagination"

type AccountDTO struct {
	Code        string `json:"code" form:"code" validate:"required,numeric"`
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description" validate:"omitempty"`
	Normal      int    `json:"normal" form:"normal" validate:"required,oneof=1 -1"`
	Category    uint   `json:"category" form:"category" validate:"required,exist=account_categories"`
	Company     uint   `json:"company" form:"company" validate:"required,exist=companies"`
}

type AccountQuery struct {
	pagination.Pagination
	Keyword  string `query:"keyword"`
	Category int    `query:"category"`
	Company  int    `query:"company"`
}
