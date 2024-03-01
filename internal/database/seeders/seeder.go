package seeders

import "gorm.io/gorm"

func DatabaseSeeder(db *gorm.DB) {
	AccountSeeder(db)
}
