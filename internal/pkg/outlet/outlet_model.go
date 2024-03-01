package outlet

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/company"
	"abude-backend/internal/pkg/employee"
)

type OutletCount struct {
	TotalCount    int `json:"totalCount"`
	ActiveCount   int `json:"activeCount"`
	InactiveCount int `json:"inactiveCount"`
}

type Outlet struct {
	common.BaseModel
	Name      string  `json:"name" gorm:"type:varchar(100)"`
	Address   string  `json:"address" gorm:"type:varchar(255)"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    bool    `json:"status"`

	Company   *company.Company `json:"company,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	CompanyID uint             `json:"-"`
}

type OutletEmployee struct {
	common.BaseModel
	Type string `json:"type" gorm:"type:enum('admin','employee')" enums:"admin,employee"`

	Employee   *employee.Employee `json:"employee,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	EmployeeID uint               `json:"-"`

	Outlet   *Outlet `json:"outlet,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint    `json:"-"`
}
