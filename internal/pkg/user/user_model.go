package user

import (
	"abude-backend/internal/common"
	"abude-backend/pkg/exception"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	RoleSuperadmin = "superadmin"
	RoleOwner      = "owner"
	RoleEmployee   = "employee"
)

type User struct {
	common.BaseModel
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Username string `json:"username" gorm:"type:varchar(50)"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Role     string `json:"role" gorm:"type:enum('superadmin','owner','employee')" enums:"superadmin,owner,employee"`
	Status   bool   `json:"status"`
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&user).Where("username = ? AND id != ?", user.Username, user.ID).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return exception.Validation(map[string]string{
			"username": "Username telah digunakan",
		})
	}

	return nil
}

func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
}

type WithEditor struct {
	Editor   *User `json:"editor,omitempty" gorm:"constraint:OnDelete:SET NULL;" swaggerignore:"true"`
	EditorID *uint `json:"-" gorm:"column:id_editor"`
}
