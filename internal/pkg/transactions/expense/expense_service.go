package expense

import (
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"
	"time"

	"gorm.io/gorm"
)

type ExpenseService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *ExpenseService {
	return &ExpenseService{db}
}

func (s *ExpenseService) FindOne(id int) (*Expense, error) {
	var expense Expense
	if err := s.db.Preload("Account").Preload("Outlet").Preload("Company").First(&expense, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &expense, nil
}

func (s *ExpenseService) FindAll(query ExpenseQuery) *pagination.Result[Expense] {
	result := pagination.New[Expense](query.Pagination)

	db := s.db.Model(&Expense{}).Preload("Account")
	if query.Account != 0 {
		db.Where("account_id = ?", query.Account)
	}

	if query.Company != 0 {
		db.Where("company_id = ?", query.Company)
	}

	if query.Outlet != 0 {
		db.Where("outlet_id = ?", query.Outlet)
	}

	db.Order("created_at DESC")

	return result.Paginate(db)
}

func (s *ExpenseService) Create(data ExpenseDTO) (*Expense, error) {
	expense := Expense{
		Amount:    data.Amount,
		Type:      data.Type,
		Status:    StatusAccepted,
		Date:      time.Now(),
		Notes:     data.Notes,
		CompanyID: data.Company,
		OutletID:  data.Outlet,
		AccountID: data.Account,
	}

	if !data.Date.IsZero() {
		expense.Date = data.Date
	}

	if err := s.db.Create(&expense).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &expense, nil
}

func (s *ExpenseService) Update(id int, data ExpenseDTO) (*Expense, error) {
	var expense Expense
	if err := s.db.First(&expense, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	expense.Amount = data.Amount
	expense.Type = data.Type
	expense.Status = StatusAccepted
	expense.Date = data.Date
	expense.Notes = data.Notes
	expense.CompanyID = data.Company
	expense.OutletID = data.Outlet
	expense.AccountID = data.Account

	if !data.Date.IsZero() {
		expense.Date = data.Date
	}

	if err := s.db.Save(&expense).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &expense, nil
}

func (s *ExpenseService) Delete(id int) (*Expense, error) {
	var expense Expense
	if err := s.db.First(&expense, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&expense).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &expense, nil
}

func (s *ExpenseService) SetStatus(id int, status string) error {
	var expense Expense
	if err := s.db.First(&expense, id).Error; err != nil {
		return exception.DB(err)
	}

	if expense.Status == status {
		return exception.BadRequest("Status tidak berubah")
	}

	expense.Status = status

	if err := s.db.Save(&expense).Error; err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *ExpenseService) GetSummary(query ExpenseSummaryQuery) ([]ExpenseSummary, error) {
	var summary []ExpenseSummary

	db := s.db.Model(&Expense{})
	db.Select("products.id, products.name, SUM(expense_items.quantity) AS quantity, SUM(expense_items.total) AS total, DATE(expenses.date) AS date")
	db.Joins("INNER JOIN outlet_expenses ON expenses.id = outlet_expenses.expense_id")
	db.Joins("RIGHT JOIN expense_items ON expenses.id = expense_items.expense_id")
	db.Joins("INNER JOIN products ON products.id = expense_items.product_id")
	db.Where("expenses.status != ?", "canceled")

	if query.StartDate != "" {
		db.Where("DATE(expenses.date) >= ?", query.StartDate)
	}

	if query.EndDate != "" {
		db.Where("DATE(expenses.date) <= ?", query.EndDate)
	}

	if len(query.Status) > 0 {
		db.Where("expenses.status IN (?)", query.Status)
	}

	if query.Outlet != 0 {
		db.Where("outlet_expenses.outlet_id = ?", query.Outlet)
	}

	db.Group("expense_items.product_id, DATE(expenses.date)")
	db.Order("DATE(expenses.date) ASC")

	if err := db.Find(&summary).Error; err != nil {
		return nil, exception.DB(err)
	}

	return summary, nil
}

func (s *ExpenseService) Using(tx *gorm.DB) *ExpenseService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *ExpenseService) WithContext(ctx context.Context) *ExpenseService {
	s.db = s.db.WithContext(ctx)

	return s
}
