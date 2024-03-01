package wage

import (
	"abude-backend/internal/pkg/accounts/account"
	"abude-backend/internal/pkg/transactions/expense"
	"abude-backend/pkg/exception"
	"abude-backend/pkg/pagination"
	"context"
	"time"

	"gorm.io/gorm"
)

type WageService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *WageService {
	return &WageService{db}
}

func (s *WageService) FindOne(id int) (*Wage, error) {
	var wage Wage
	if err := s.db.Preload("Employee").Preload("Outlet").Preload("Company").First(&wage, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &wage, nil
}

func (s *WageService) FindAll(query WageQuery) *pagination.Result[Wage] {
	result := pagination.New[Wage](query.Pagination)

	db := s.db.Model(&Wage{}).Preload("Employee")
	if query.Employee != 0 {
		db.Where("employee_id = ?", query.Employee)
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

func (s *WageService) Create(data WageDTO) (*Wage, error) {
	wage := Wage{
		Amount:     data.Amount,
		Type:       data.Type,
		Status:     StatusAccepted,
		Date:       time.Now(),
		Notes:      data.Notes,
		CompanyID:  data.Company,
		OutletID:   data.Outlet,
		EmployeeID: data.Employee,
	}

	if !data.Date.IsZero() {
		wage.Date = data.Date
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		expenseService := expense.NewService(tx)

		var acc account.Account
		if err := tx.Where("code = ? AND (company_id = ? OR company_id IS NULL)", account.CodeExpenseWage, data.Company).First(&acc).Error; err != nil {
			return err
		}

		if err := tx.Create(&wage).Error; err != nil {
			return err
		}

		expenseService.Create(expense.ExpenseDTO{
			Amount:  wage.Amount,
			Type:    wage.Type,
			Date:    data.Date,
			Notes:   data.Notes,
			Account: acc.ID,
			Company: data.Company,
			Outlet:  data.Outlet,
		})

		return nil
	})

	if err != nil {
		return nil, exception.DB(err)
	}

	return &wage, nil
}

func (s *WageService) Update(id int, data WageDTO) (*Wage, error) {
	var wage Wage
	if err := s.db.First(&wage, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	wage.Amount = data.Amount
	wage.Type = data.Type
	wage.Status = StatusAccepted
	wage.Date = data.Date
	wage.Notes = data.Notes
	wage.CompanyID = data.Company
	wage.OutletID = data.Outlet
	wage.EmployeeID = data.Employee

	if !data.Date.IsZero() {
		wage.Date = data.Date
	}

	if err := s.db.Save(&wage).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &wage, nil
}

func (s *WageService) Delete(id int) (*Wage, error) {
	var wage Wage
	if err := s.db.First(&wage, id).Error; err != nil {
		return nil, exception.DB(err)
	}

	if err := s.db.Delete(&wage).Error; err != nil {
		return nil, exception.DB(err)
	}

	return &wage, nil
}

func (s *WageService) SetStatus(id int, status string) error {
	var wage Wage
	if err := s.db.First(&wage, id).Error; err != nil {
		return exception.DB(err)
	}

	if wage.Status == status {
		return exception.BadRequest("Status tidak berubah")
	}

	wage.Status = status

	if err := s.db.Save(&wage).Error; err != nil {
		return exception.DB(err)
	}

	return nil
}

func (s *WageService) GetSummary(query WageSummaryQuery) ([]WageSummary, error) {
	var summary []WageSummary

	db := s.db.Model(&Wage{})
	db.Select("products.id, products.name, SUM(wage_items.quantity) AS quantity, SUM(wage_items.total) AS total, DATE(wages.date) AS date")
	db.Joins("INNER JOIN outlet_wages ON wages.id = outlet_wages.wage_id")
	db.Joins("RIGHT JOIN wage_items ON wages.id = wage_items.wage_id")
	db.Joins("INNER JOIN products ON products.id = wage_items.product_id")
	db.Where("wages.status != ?", "canceled")

	if query.StartDate != "" {
		db.Where("DATE(wages.date) >= ?", query.StartDate)
	}

	if query.EndDate != "" {
		db.Where("DATE(wages.date) <= ?", query.EndDate)
	}

	if len(query.Status) > 0 {
		db.Where("wages.status IN (?)", query.Status)
	}

	if query.Outlet != 0 {
		db.Where("outlet_wages.outlet_id = ?", query.Outlet)
	}

	db.Group("wage_items.product_id, DATE(wages.date)")
	db.Order("DATE(wages.date) ASC")

	if err := db.Find(&summary).Error; err != nil {
		return nil, exception.DB(err)
	}

	return summary, nil
}

func (s *WageService) Using(tx *gorm.DB) *WageService {
	db := s.db

	defer func() {
		s.db = db
	}()

	s.db = tx.WithContext(tx.Statement.Context)

	return s
}

func (s *WageService) WithContext(ctx context.Context) *WageService {
	s.db = s.db.WithContext(ctx)

	return s
}
