package handover

import (
	"abude-backend/internal/pkg/transactions/expense"
	"abude-backend/internal/pkg/transactions/purchase"
	"abude-backend/internal/pkg/transactions/sale"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"

	"gorm.io/gorm"
)

type HandoverService struct {
	db *gorm.DB
}

func NewHandoverService(db *gorm.DB) *HandoverService {
	return &HandoverService{db}
}

func (s *HandoverService) FindOne(id int) (*Handover, error) {
	var handover Handover
	if err := s.db.First(&handover, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Table("sale_items").
		Select("SUM(sale_items.price) AS price, SUM(sale_items.quantity) AS quantity, SUM(sale_items.total) AS total, sale_items.product_id as product_id, DATE(sales.date) as date").
		Joins("INNER JOIN sales ON sales.id = sale_items.sale_id").
		Where("sale_id IN (?)", s.db.
			Table("handover_sales").
			Select("sale_id").
			Where("handover_id = ?", id),
		).Preload("Product").Group("DATE(sales.date), sale_items.product_id").
		Order("DATE(sales.date) ASC").
		Find(&handover.SaleItems).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Table("purchase_items").
		Select("SUM(purchase_items.price) AS price, SUM(purchase_items.quantity) AS quantity, SUM(purchase_items.total) AS total, purchase_items.product_id as product_id, DATE(purchases.date) as date").
		Joins("INNER JOIN purchases ON purchases.id = purchase_items.purchase_id").
		Where("purchase_id IN (?)", s.db.
			Table("handover_purchases").
			Select("purchase_id").
			Where("handover_id = ?", id),
		).Preload("Product").Group("DATE(purchases.date), purchase_items.product_id").
		Order("DATE(purchases.date) ASC").
		Find(&handover.PurchaseItems).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &handover, nil
}

func (s *HandoverService) FindAll(query HandoverQuery) *pagination.Result[Handover] {
	result := pagination.New[Handover](query.Pagination)

	db := s.db.Model(&Handover{}).Preload("Shift").Preload("Outlet").Preload("Editor")
	if query.Outlet != 0 {
		db.Where("outlet_id = ?", query.Outlet)
	}

	if query.Shift != 0 {
		db.Where("shift_id = ?", query.Shift)
	}

	if !query.StartDate.IsZero() {
		db.Where("date >= ?", query.StartDate)
	}

	if !query.EndDate.IsZero() {
		db.Where("date <= ?", query.EndDate)
	}

	db.Order("date DESC")

	return result.Paginate(db)
}

func (s *HandoverService) Create(data HandoverDTO) (*Handover, error) {
	handover := Handover{
		Note:         data.Note,
		CashReceived: data.CashReceived,
		CashReturned: data.CashReturned,
		Date:         data.Date,
		OutletID:     data.Outlet,
		ShiftID:      data.Shift,
	}

	sales := sale.NewService(s.db).FindAll(sale.SaleQuery{
		Outlet: data.Outlet,
		Status: []string{sale.StatusAccepted},
		Pagination: pagination.Pagination{
			Limit: -1,
		},
	})

	purchases := purchase.NewService(s.db).FindAll(purchase.PurchaseQuery{
		Outlet: data.Outlet,
		Status: []string{purchase.StatusAccepted},
		Pagination: pagination.Pagination{
			Limit: -1,
		},
	})

	expenses := expense.NewService(s.db).FindAll(expense.ExpenseQuery{
		Outlet: data.Outlet,
		Status: []string{expense.StatusAccepted},
		Pagination: pagination.Pagination{
			Limit: -1,
		},
	})

	var saleIds []uint
	for _, sale := range sales.Result {
		saleIds = append(saleIds, sale.ID)
		handover.SalesTotal += sale.Total
	}

	var purchaseIds []uint
	for _, purchase := range purchases.Result {
		purchaseIds = append(purchaseIds, purchase.ID)
		handover.PurchasesTotal += purchase.Total
	}

	var expenseIds []uint
	for _, expense := range expenses.Result {
		expenseIds = append(expenseIds, expense.ID)
		handover.ExpensesTotal += expense.Amount
	}

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&handover).Error; err != nil {
			return err
		}

		if err := tx.Model(&handover).Association("Sales").Append(&sales.Result); err != nil {
			return err
		}

		if err := tx.Model(&handover).Association("Purchases").Append(&purchases.Result); err != nil {
			return err
		}

		if err := tx.Model(&handover).Association("Expenses").Append(&expenses.Result); err != nil {
			return err
		}

		if err := tx.Model(&sale.Sale{}).Where("id IN (?)", saleIds).Update("status", sale.StatusApproved).Error; err != nil {
			return err
		}

		if err := tx.Model(&purchase.Purchase{}).Where("id IN (?)", purchaseIds).Update("status", purchase.StatusApproved).Error; err != nil {
			return err
		}

		if err := tx.Model(&expense.Expense{}).Where("id IN (?)", expenseIds).Update("status", expense.StatusApproved).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, exception.DB(err)
	}

	return &handover, nil
}

func (s *HandoverService) Update(id int, data interface{}) (*Handover, error) {
	var handover Handover
	if err := s.db.First(&handover, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Save(&handover).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &handover, nil
}

func (s *HandoverService) Delete(id int) (*Handover, error) {
	var handover Handover
	if err := s.db.First(&handover, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&handover).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &handover, nil
}

func (s *HandoverService) Using(tx *gorm.DB) *HandoverService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *HandoverService) WithContext(ctx context.Context) *HandoverService {
	s.db = s.db.WithContext(ctx)

	return s
}
