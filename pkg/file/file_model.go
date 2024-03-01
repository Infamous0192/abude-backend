package file

import (
	"abude-backend/pkg/pagination"
	"mime/multipart"
	"time"
)

type File struct {
	ID           uint      `json:"id" gorm:"primarykey, autoIncrement"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"originalname"`
	Path         string    `json:"path"`
	Extension    string    `json:"extension"`
	Size         int       `json:"size"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}

type FileUploadRequest struct {
	File *multipart.FileHeader `form:"file" validate:"required" swaggerignore:"true"`
}

type FileQuery struct {
	pagination.Pagination
}
