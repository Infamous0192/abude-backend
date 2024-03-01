package handover

import (
	"abude-backend/internal/common"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/user"
	"time"
)

type Proof struct {
	common.BaseModel
	user.WithEditor
	Date     time.Time `json:"date"`
	Evidence string    `json:"evidence"`

	Outlet   *outlet.Outlet `json:"outlet,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	OutletID uint           `json:"-"`
}

func (Proof) TableName() string {
	return "handover_proofs"
}
