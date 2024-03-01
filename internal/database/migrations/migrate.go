package migrations

import (
	"abude-backend/internal/pkg/attendances/shift"
	"abude-backend/internal/pkg/company"
	"abude-backend/internal/pkg/employee"
	"abude-backend/internal/pkg/handover"
	"abude-backend/internal/pkg/inventories/category"
	"abude-backend/internal/pkg/inventories/product"
	"abude-backend/internal/pkg/outlet"
	"abude-backend/internal/pkg/supplier"
	"abude-backend/internal/pkg/transactions/expense"
	"abude-backend/internal/pkg/transactions/purchase"
	"abude-backend/internal/pkg/transactions/sale"
	"abude-backend/internal/pkg/transactions/wage"
	"abude-backend/internal/pkg/turnover"
	"abude-backend/internal/pkg/user"
	"errors"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&company.Company{},
		&employee.Employee{},
		&outlet.Outlet{},
		&outlet.OutletEmployee{},
		&category.Category{},
		&product.Ingredient{},
		&product.Product{},
		&supplier.Supplier{},
		&sale.Sale{},
		&sale.SaleItem{},
		&sale.OutletSale{},
		&purchase.Purchase{},
		&purchase.PurchaseItem{},
		&purchase.OutletPurchase{},
		&expense.Expense{},
		&wage.Wage{},
		&handover.Handover{},
		&handover.Proof{},
		&turnover.Turnover{},
		&shift.Shift{},
	)

	MigrateAccount(db)
	MigrateInventory(db)

	if err := db.First(&user.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		user := user.User{
			Name:     "Tarkiz Paz Banua",
			Username: "tpaztech",
			Role:     user.RoleSuperadmin,
			Status:   true,
		}

		user.SetPassword("asdqwe123")

		db.Create(&user)
	}
}
