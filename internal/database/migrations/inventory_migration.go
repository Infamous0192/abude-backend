package migrations

import (
	"abude-backend/internal/pkg/inventories/category"
	"abude-backend/internal/pkg/inventories/inventory"
	"abude-backend/internal/pkg/inventories/product"

	"gorm.io/gorm"
)

func MigrateInventory(db *gorm.DB) {
	err := db.AutoMigrate(
		&category.Category{},
		&product.Product{},
		&product.Ingredient{},
		&inventory.Inventory{},
		&inventory.OutletInventory{},
		&inventory.Recapitulation{},
		&inventory.RecapitulationItem{},
	)

	if err != nil {
		panic(err)
	}
}
