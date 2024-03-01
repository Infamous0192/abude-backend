package migrations

import (
	"abude-backend/internal/pkg/accounts/account"
	"abude-backend/internal/pkg/accounts/category"

	"gorm.io/gorm"
)

func MigrateAccount(db *gorm.DB) {
	err := db.AutoMigrate(
		&account.Account{},
		&category.Category{},
	)

	if err != nil {
		panic(err)
	}
}
