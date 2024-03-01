package common

import "time"

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey, autoIncrement"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}
